package config

import (
	"path/filepath"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/pixelfs/pixelfs/util"
	"github.com/spf13/viper"
)

type Config struct {
	Endpoint string `toml:"endpoint"`
	Token    string `toml:"token"`
	Pwd      string `toml:"pwd"`
	Debug    bool   `toml:"debug"`
	Daemon   Daemon `toml:"daemon"`
	FFmpeg   FFmpeg `toml:"ffmpeg"`
	Webdav   Webdav `toml:"webdav"`
}

type Daemon struct {
	Listen string `toml:"listen"`
}

type Log struct {
	Level string `toml:"level"`
}

type Cache struct {
	Path   string `toml:"path"`
	Expire int    `toml:"expire"`
}

type FFmpeg struct {
	Cache   Cache          `toml:"cache"`
	Options map[string]any `toml:"options"`
}

type CORS struct {
	Credentials   bool     `toml:"credentials"`
	AllowOrigin   []string `toml:"allow_origin"`
	AllowHeaders  []string `toml:"allow_headers"`
	AllowMethods  []string `toml:"allow_methods"`
	ExposeHeaders []string `toml:"expose_headers"`
	MaxAge        int      `toml:"max_age"`
}

type Webdav struct {
	Listen string `toml:"listen"`
	Prefix string `toml:"prefix"`
	Cache  Cache  `toml:"cache"`
	CORS   CORS   `toml:"cors"`
	Users  []User `toml:"users"`
}

func LoadConfig(path string, isFile bool) error {
	home, err := util.GetHomeDir()
	if err != nil {
		return err
	}

	if isFile {
		viper.SetConfigFile(path)
	} else {
		viper.SetConfigName("config")
		if path == "" {
			viper.AddConfigPath(home)
		} else {
			viper.AddConfigPath(path)
		}
	}

	viper.SetConfigType("toml")

	viper.SetDefault("endpoint", "https://pixelfs.io")
	viper.SetDefault("daemon.listen", "0.0.0.0:15233")

	// ffmpeg
	viper.SetDefault("ffmpeg.cache.path", filepath.Join(home, "cache", "ffmpeg"))
	viper.SetDefault("ffmpeg.cache.expire", 3600*24)

	// webdav
	viper.SetDefault("webdav.listen", "0.0.0.0:5233")
	viper.SetDefault("webdav.cache.path", filepath.Join(home, "cache", "webdav"))
	viper.SetDefault("webdav.cache.expire", 3600*24)
	viper.SetDefault("webdav.cors.allow_origin", []string{"*"})
	viper.SetDefault("webdav.cors.allow_headers", []string{"*"})
	viper.SetDefault("webdav.cors.allow_methods", []string{"*"})

	viper.SetEnvPrefix("pixelfs")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			if err := viper.SafeWriteConfig(); err != nil {
				return err
			}
		}
	}

	return nil
}

func Set(key string, v any) error {
	viper.Set(key, v)
	return viper.WriteConfig()
}

func Remove(key string) error {
	switch viper.Get(key).(type) {
	case string:
		viper.Set(key, "")
	default:
		viper.Set(key, nil)
	}

	return viper.WriteConfig()
}

func GetConfig() (*Config, error) {
	cfg := &Config{}
	err := viper.Unmarshal(cfg, viper.DecodeHook(
		mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToSliceHookFunc(","),
			mapstructure.TextUnmarshallerHookFunc(),
		),
	))

	if err != nil {
		return nil, err
	}

	return cfg, nil
}
