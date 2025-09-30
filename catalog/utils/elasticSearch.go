package utils

import (
	"context"
	"fmt"
	"gopkg.in/olivere/elastic.v5"
	"time"
)

type ElasticSearch struct {
	Host     string
	Port     string
	Username string
	Password string
	Timeout  time.Duration
}

type Options func(*ElasticSearch)

func WithHost(host string) Options {
	return func(e *ElasticSearch) {
		e.Host = host
	}
}

func WithPort(port string) Options {
	return func(e *ElasticSearch) {
		e.Port = port
	}
}

func WithUsername(username string) Options {
	return func(e *ElasticSearch) {
		e.Username = username
	}
}

func WithPassword(password string) Options {
	return func(e *ElasticSearch) {
		e.Password = password
	}
}

func WithTimeout(timeout time.Duration) Options {
	return func(e *ElasticSearch) {
		e.Timeout = timeout
	}
}

func (e *ElasticSearch) uri() string {
	return fmt.Sprintf("http://%s:%s", e.Host, e.Port)
}

func (e *ElasticSearch) Connect() (*elastic.Client, error) {
	client, err := elastic.NewClient(
		elastic.SetURL(e.uri()),
		elastic.SetBasicAuth(e.Username, e.Password),
		elastic.SetHealthcheckTimeoutStartup(e.Timeout),
		elastic.SetSniff(false),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create elasticsearch client: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), e.Timeout)
	defer cancel()

	_, _, err = client.Ping(e.uri()).Do(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to ping elasticsearch client: %w", err)
	}

	return client, nil
}

func NewElasticSearch(opts ...Options) *ElasticSearch {
	es := &ElasticSearch{
		Timeout: 5 * time.Second,
	}
	for _, opt := range opts {
		opt(es)
	}
	return es
}
