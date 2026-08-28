package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port            string `envconfig:"PORT" default:"8080"`
	APIKey          string `envconfig:"GEMINI_API_KEY" required:"true"`
	PathToInterview string `envconfig:"PATH_TO_INTERVIEW" default:"../interviews/"`
	PathToChats     string `envconfig:"PATH_TO_CHATS" default:"../chats/"`
	DefaultModel    string `envconfig:"GEMINI_DEFAULT_MODEL" default:"gemini-3.1-flash-lite"`
}

func newConfig() (*Config, error) {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, fmt.Errorf("process envconfig: %w", err)
	}
	return &cfg, nil
}

// NewConfigMust собирает конфигурацию и паникует при ошибке.
// Нет смысла в запуске приложения, если есть ошибки в конфиге (паттерн Must).
func NewConfigMust() *Config {
	cfg, err := newConfig()
	if err != nil {
		panic(fmt.Errorf("get FS config: %w", err))
	}
	return cfg
}
