package config

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

type User struct {
	Username    string      `toml:"username"`
	Password    string      `toml:"password"`
	Permissions Permissions `toml:"permissions"`
	Rules       []*Rule     `toml:"rules"`
}

func (u User) checkPassword(input string) bool {
	return u.Password == input
}

func (u User) Allowed(r *http.Request, fileExists func(string) bool) bool {
	if r.Method == "COPY" || r.Method == "MOVE" {
		dst := r.Header.Get("Destination")

		for i := len(u.Rules) - 1; i >= 0; i-- {
			if u.Rules[i].Matches(dst) {
				if !u.Rules[i].Permissions.AllowedDestination(r, fileExists) {
					return false
				}
				break
			}
		}

		if !u.Permissions.AllowedDestination(r, fileExists) {
			return false
		}
	}

	for i := len(u.Rules) - 1; i >= 0; i-- {
		if u.Rules[i].Matches(r.URL.Path) {
			return u.Rules[i].Permissions.Allowed(r, fileExists)
		}
	}

	return u.Permissions.Allowed(r, fileExists)
}

func (u *User) Validate() error {
	for _, r := range u.Rules {
		if err := r.Validate(); err != nil {
			return fmt.Errorf("invalid permissions: %w", err)
		}
	}

	return nil
}

type Rule struct {
	Permissions Permissions    `toml:"permissions"`
	Path        string         `toml:"path"`
	Regex       *regexp.Regexp `toml:"regex"`
}

func (r *Rule) Validate() error {
	if r.Regex == nil && r.Path == "" {
		return errors.New("invalid rule: must either define a path of a regex")
	}

	if r.Regex != nil && r.Path != "" {
		return errors.New("invalid rule: cannot define both regex and path")
	}

	return nil
}

func (r *Rule) Matches(path string) bool {
	if r.Regex != nil {
		return r.Regex.MatchString(path)
	}

	return strings.HasPrefix(path, r.Path)
}

type Permissions struct {
	Create bool
	Read   bool
	Update bool
	Delete bool
}

func (p *Permissions) UnmarshalText(data []byte) error {
	text := strings.ToLower(string(data))
	if text == "none" {
		return nil
	}

	for _, c := range text {
		switch c {
		case 'c':
			p.Create = true
		case 'r':
			p.Read = true
		case 'u':
			p.Update = true
		case 'd':
			p.Delete = true
		default:
			return fmt.Errorf("invalid permission: %q", c)
		}
	}

	return nil
}

func (p Permissions) Allowed(r *http.Request, fileExists func(string) bool) bool {
	switch r.Method {
	case "GET", "HEAD", "OPTIONS", "POST", "PROPFIND":
		return p.Read
	case "MKCOL":
		return p.Create
	case "PROPPATCH":
		return p.Update
	case "PUT":
		if fileExists(r.URL.Path) {
			return p.Update
		} else {
			return p.Create
		}
	case "COPY":
		return p.Read
	case "MOVE":
		return p.Read && p.Delete
	case "DELETE":
		return p.Delete
	case "LOCK", "UNLOCK":
		return p.Create || p.Read || p.Update || p.Delete
	default:
		return false
	}
}

func (p Permissions) AllowedDestination(r *http.Request, fileExists func(string) bool) bool {
	switch r.Method {
	case "COPY", "MOVE":
		if fileExists(r.Header.Get("Destination")) {
			return p.Update
		} else {
			return p.Create
		}
	default:
		return false
	}
}
