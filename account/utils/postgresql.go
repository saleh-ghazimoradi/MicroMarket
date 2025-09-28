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

type Options func(*Postgresql)

func WithHost(host string) Options {
	return func(p *Postgresql) {
		p.Host = host
	}
}

func WithPort(port string) Options {
	return func(p *Postgresql) {
		p.Port = port
	}
}

func WithUser(user string) Options {
	return func(p *Postgresql) {
		p.User = user
	}
}

func WithPassword(password string) Options {
	return func(p *Postgresql) {
		p.Password = password
	}
}

func WithName(name string) Options {
	return func(p *Postgresql) {
		p.Name = name
	}
}

func WithMaxOpenConn(maxConn int) Options {
	return func(p *Postgresql) {
		p.MaxOpenConn = maxConn
	}
}

func WithMaxIdleConn(maxIdleConn int) Options {
	return func(p *Postgresql) {
		p.MaxIdleConn = maxIdleConn
	}
}

func WithMaxIdleTime(maxIdleTime time.Duration) Options {
	return func(p *Postgresql) {
		p.MaxIdleTime = maxIdleTime
	}
}

func WithSSLMode(mode string) Options {
	return func(p *Postgresql) {
		p.SSLMode = mode
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

func NewPostgresql(opts ...Options) *Postgresql {
	p := &Postgresql{}
	for _, opt := range opts {
		opt(p)
	}
	return p
}
