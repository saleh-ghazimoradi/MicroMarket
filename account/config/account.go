package config

type AccountServer struct {
	GRPCPort    string `env:"GRPC_PORT"`
	AccountURL  string `env:"ACCOUNT_URL"`
	GraphqlPort string `env:"GRAPHQL_PORT"`
}
