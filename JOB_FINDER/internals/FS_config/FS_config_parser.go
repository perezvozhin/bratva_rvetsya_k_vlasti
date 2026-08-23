package FS_config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	PathFilesystem string `envconfig:"PATH_TO_FILESYSTEM" required:"true"`
	ApiKey         string `envconfig:"API_KEY" required:"true"`
	Port           string `envconfig:"PORT" default:"9090"`
	Name           string `envconfig:"NAME" default:"JOB_FINDER"`
	TimeZone       string `envconfig:"TIMEZONE" default:"UTC"`
}

func newConfig() (*Config, error) {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, fmt.Errorf("process envconfig: %w", err)
	}
	return &cfg, nil
}

// Init собирает конфигурацию и паникует при ошибке.
// Нет смысла в запуске приложения, если есть ошибки в конфиге (паттерн Must).
func Init() *Config {
	cfg, err := newConfig()
	if err != nil {
		panic(fmt.Errorf("get FS config: %w", err))
	}
	return cfg
}
