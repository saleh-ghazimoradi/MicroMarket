package main

import "github.com/99designs/gqlgen/graphql"

type GraphqlServer struct {
	//accountClient *account.Client
	//catalogClient *catalog.Client
	//orderClient   *orderClient
}

func NewGraphqlServer(accountURL, catalogURL, orderURL string) (*GraphqlServer, error) {
	//accountClient, err := account.NewClient(accountURL)
	//if err != nil {
	//	return nil, err
	//}
	//
	//catalogClient, err := catalog.NewClient(catalogURL)
	//if err != nil {
	//	accountClient.Close()
	//	return nil, err
	//}
	//
	//orderClient, err := order.NewClient(orderURL)
	//if err != nil {
	//	accountClient.Close()
	//	catalogClient.Close()
	//	return nil, err
	//}

	return &GraphqlServer{
		//accountClient,
		//catalogClient,
		//orderClient,
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
