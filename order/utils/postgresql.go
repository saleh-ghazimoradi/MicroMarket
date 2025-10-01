package utils

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type Postgresql struct {
	Host        string
	Port        string
	User        string
	Password    string
	Name        string
	MaxOpenConn int
	MaxIdleConn int
	MaxIdleTime time.Duration
	SSLMode     string
}

type Option func(*Postgresql)

func WithHost(host string) Option {
	return func(o *Postgresql) {
		o.Host = host
	}
}

func WithPort(port string) Option {
	return func(o *Postgresql) {
		o.Port = port
	}
}

func WithUser(user string) Option {
	return func(o *Postgresql) {
		o.User = user
	}
}

func WithPassword(password string) Option {
	return func(o *Postgresql) {
		o.Password = password
	}
}

func WithName(name string) Option {
	return func(o *Postgresql) {
		o.Name = name
	}
}

func WithMaxOpenConn(maxOpenConn int) Option {
	return func(o *Postgresql) {
		o.MaxOpenConn = maxOpenConn
	}
}

func WithMaxIdleConn(maxIdleConn int) Option {
	return func(o *Postgresql) {
		o.MaxIdleConn = maxIdleConn
	}
}

func WithMaxIdleTime(maxIdleTime time.Duration) Option {
	return func(o *Postgresql) {
		o.MaxIdleTime = maxIdleTime
	}
}

func WithSSLMode(mode string) Option {
	return func(o *Postgresql) {
		o.SSLMode = mode
	}
}

func (p *Postgresql) uri() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", p.Host, p.Port, p.User, p.Password, p.Name, p.SSLMode)
}

func (p *Postgresql) Connect() (*sql.DB, error) {
	db, err := sql.Open("postgres", p.uri())
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(p.MaxOpenConn)
	db.SetMaxIdleConns(p.MaxIdleConn)
	db.SetConnMaxLifetime(p.MaxIdleTime)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = db.PingContext(ctx); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

func NewPostgresql(opts ...Option) *Postgresql {
	p := &Postgresql{}
	for _, opt := range opts {
		opt(p)
	}
	return p
}
