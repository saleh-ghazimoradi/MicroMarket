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
	Order      Order
	Postgresql Postgresql
}

func NewConfig() (*Config, error) {
	once.Do(func() {
		instance = &Config{}
		initErr = env.Parse(instance)
		if initErr != nil {
			initErr = nil
		}
	})
	return instance, initErr
}
