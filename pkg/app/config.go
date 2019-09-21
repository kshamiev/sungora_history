package app

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

const configYaml = "config.yaml"

// ConfigSearchPath поиск конфигурации
func ConfigSearchPath(serviceName string) (string, error) {
	sep := string(os.PathSeparator)
	path := filepath.Dir(filepath.Dir(os.Args[0]))
	if path == "." {
		path, _ = os.Getwd()
		path = filepath.Dir(path)
	}
	path += sep + configYaml
	//
	if inf, err := os.Stat(path); err == nil {
		if inf.Mode().IsRegular() {
			return path, nil
		}
	}
	return "", errors.New("config file not found")
}

// ConfigLoad загрузка конфигурации
func ConfigLoad(path string, cfg interface{}) (err error) {
	var data []byte
	if data, err = ioutil.ReadFile(path); err != nil {
		return
	}
	return yaml.Unmarshal(data, cfg)
}

// ConfigSetDefault инициализация дефолтовыми значениями
func ConfigSetDefault(cfg *ConfigApp) {
	if cfg == nil {
		cfg = &ConfigApp{}
	}
	// режим работы приложения
	if cfg.Mode == "" {
		cfg.Mode = "dev"
	}
	// пути
	sep := string(os.PathSeparator)
	if cfg.DirWork == "" {
		cfg.DirWork, _ = os.Getwd()
		sl := strings.Split(cfg.DirWork, sep)
		if sl[len(sl)-1] == "bin" {
			sl = sl[:len(sl)-1]
		}
		cfg.DirWork = strings.Join(sl, sep)
	}
	// сессия
	if cfg.SessionTimeout == 0 {
		cfg.SessionTimeout = time.Duration(14400) * time.Second
	}
	//
	if cfg.Domain == "" {
		cfg.Domain = "localhost"
	}
}
