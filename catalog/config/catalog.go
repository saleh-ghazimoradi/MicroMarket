package config

type CatalogServer struct {
	Port string `env:"CATALOG_SERVER_PORT"`
}
