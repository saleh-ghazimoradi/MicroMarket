package config

import "time"

type Postgresql struct {
	Host        string        `env:"POSTGRESQL_HOST"`
	Port        string        `env:"POSTGRESQL_PORT"`
	User        string        `env:"POSTGRESQL_USER"`
	Password    string        `env:"POSTGRESQL_PASSWORD"`
	Name        string        `env:"POSTGRESQL_NAME"`
	MaxOpenConn int           `env:"POSTGRESQL_MAX_OPEN_CONN"`
	MaxIdleConn int           `env:"POSTGRESQL_MAX_IDLE_CONN"`
	MaxIdleTime time.Duration `env:"POSTGRESQL_MAX_IDLE_TIME"`
	SSLMode     string        `env:"POSTGRESQL_SSL_MODE"`
}
