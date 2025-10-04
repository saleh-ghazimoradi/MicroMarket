package config

type Order struct {
	AccountURL string `env:"ACCOUNT_URL"`
	CatalogURL string `env:"CATALOG_URL"`
}
