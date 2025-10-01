package main

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/saleh-ghazimoradi/MicroMarket/account/gateway/grpcAccountHandler"
	"github.com/saleh-ghazimoradi/MicroMarket/catalog/gateway/gRPCCatalogHandler"
)

type GraphqlServer struct {
	gRPCAccountClient *grpcAccountHandler.GRPCAccountClient
	gRPCCatalogClient *gRPCCatalogHandler.GRPCCatalogClient
}

func NewGraphqlServer(accountURL, catalogURL, orderURL string) (*GraphqlServer, error) {
	accountClient, err := grpcAccountHandler.NewGRPCAccountClient(accountURL)
	if err != nil {
		return nil, err
	}

	catalogClient, err := gRPCCatalogHandler.NewGrpcCatalogClient(catalogURL)
	if err != nil {
		return nil, err
	}

	return &GraphqlServer{
		gRPCAccountClient: accountClient,
		gRPCCatalogClient: catalogClient,
	}, nil

}

func (s *GraphqlServer) Query() QueryResolver {
	return &queryResolver{
		server: s,
	}
}

func (s *GraphqlServer) Mutation() MutationResolver {
	return &mutationResolver{
		server: s,
	}
}

func (s *GraphqlServer) Account() AccountResolver {
	return &accountResolver{
		server: s,
	}
}

func (s *GraphqlServer) ToExecutableSchema() graphql.ExecutableSchema {
	return NewExecutableSchema(Config{Resolvers: s})
}
