package main

import "context"

type accountResolver struct {
	server *Server
}

func (a *accountResolver) Orders(ctx context.Context, account *Account) ([]*Order, error) {
	return nil, nil
}
