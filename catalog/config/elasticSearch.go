package config

import "time"

type ElasticSearch struct {
	Host     string        `env:"elasticSearchHost"`
	Port     string        `env:"elasticSearchPort"`
	Username string        `env:"elasticSearchUsername"`
	Password string        `env:"elasticSearchPassword"`
	Timeout  time.Duration `env:"elasticSearchTimeout"`
}
