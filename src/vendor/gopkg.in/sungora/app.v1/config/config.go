package config

import (
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Name        string
	DisplayName string
	Description string
	Server      Server
	Database    Database
	Log         Log
}

type Server struct {
	Host string
	Port int
}

type Database struct {
	Host string
	Port int
}

type Log struct {
	Info        bool
	Warning     bool
	Error       bool
	OutFile     bool
	OutFilePath string
	OutStd      bool
	Traces      bool
}

// ReadToml Функция читает конфигурационный файл в формате toml. Отдельный конфиг не связанный с beego.
func ReadToml(configFile string) (cfg *Config, err error) {
	if configFile == "" {
		if dir, err := os.Getwd(); err == nil {
			configFile = dir + "/" + filepath.Base(os.Args[0]) + ".toml"
		}

	}
	if _, err := toml.DecodeFile(configFile, &cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
