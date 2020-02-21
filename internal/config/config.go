package config

import (
	"github.com/kshamiev/sungora/pkg/app"
	"github.com/kshamiev/sungora/pkg/logger"
)

type Config struct {
	App        app.ConfigApp      `yaml:"app"`
	Lg         logger.Config      `yaml:"lg"`
	ServHTTP   app.ConfigServer   `yaml:"http"`
	GRPCServer app.ConfigGRPC     `yaml:"grpcServer"`
	Cors       app.ConfigCors     `yaml:"cors"`
	Postgresql app.ConfigPostgres `yaml:"postgresql"`
}

// Get загрузка конфигурации приложения
func Get(fileConfig ...string) (cfg *Config, err error) {
	cfg = &Config{}
	for i := range fileConfig {
		if err = app.ConfigLoad(fileConfig[i], cfg); err == nil {
			app.ConfigSetDefault(&cfg.App)
			return
		}
	}
	return
}
