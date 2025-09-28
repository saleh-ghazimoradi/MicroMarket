package main

import "context"

type accountResolver struct {
	server *GraphqlServer
}

func (a *accountResolver) Orders(ctx context.Context, account *Account) ([]*Order, error) {
	return nil, nil
}
