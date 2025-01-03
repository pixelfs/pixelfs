package util

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net"
	"path/filepath"
	"strings"
	"time"

	"github.com/denisbrodbeck/machineid"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

var (
	Resty = resty.New().SetTimeout(5 * time.Minute)
)

func GetNodeId(u string) (string, error) {
	id, err := machineid.ID()
	if err != nil || id == "" {
		id, err = generateMachineId()
		if err != nil {
			return "", err
		}
	}

	mac := hmac.New(sha1.New, []byte(id))
	mac.Write([]byte(u))
	return fmt.Sprintf("0x%x", mac.Sum(nil)), nil
}

func generateMachineId() (string, error) {
	home, err := GetHomeDir()
	if err != nil {
		return "", err
	}

	machineIdPath := filepath.Join(home, "machine-id")
	data, err := ioutil.ReadFile(machineIdPath)
	if err == nil && len(data) > 0 {
		return string(data), nil
	}

	id := uuid.New().String()
	if err := ioutil.WriteFile(machineIdPath, []byte(id), 0600); err != nil {
		return "", fmt.Errorf("failed to write machine ID to file: %w", err)
	}

	return id, nil
}

func IsNodeId(id string) bool {
	return strings.HasPrefix(id, "0x") && len(id) == 42
}

func GenerateAuthToken() (string, error) {
	bytes := make([]byte, 32)

	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(bytes), nil
}

func PadLeft(str string, length int, pad string) string {
	for len(str) < length {
		str = pad + str
	}
	return str
}

func PadRight(str string, length int, pad string) string {
	for len(str) < length {
		str = str + pad
	}
	return str
}

func IsImage(name string) bool {
	ext := filepath.Ext(name)
	return ext == ".jpg" || ext == ".jpeg" || ext == ".png" || ext == ".gif" || ext == ".bmp"
}

func IsDocument(name string) bool {
	ext := filepath.Ext(name)
	return ext == ".txt" || ext == ".md" || ext == ".doc" || ext == ".docx" || ext == ".pdf"
}

func IsAvailableAddress(address string) bool {
	conn, err := net.DialTimeout("tcp", address, 3*time.Second)
	if err != nil {
		return false
	}

	conn.Close()
	return true
}
