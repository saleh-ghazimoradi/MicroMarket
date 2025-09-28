package config

import (
	"github.com/caarlos0/env/v11"
	"sync"
)

var (
	instance *Config
	once     sync.Once
	initErr  error
)

type Config struct {
	Application Application
	Server      Server
}

type Application struct {
	AccountURL string `env:"ACCOUNT_URL"`
	CatalogURL string `env:"CATALOG_URL"`
	OrderURL   string `env:"ORDER_URL"`
}

type Server struct {
	Port string `env:"SERVER_PORT"`
}

func NewConfig() (*Config, error) {
	once.Do(func() {
		instance = &Config{}
		initErr = env.Parse(instance)
		if initErr != nil {
			instance = nil
		}
	})
	return instance, initErr
}
