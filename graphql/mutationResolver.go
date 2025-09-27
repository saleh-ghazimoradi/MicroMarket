package main

import "context"

type mutationResolver struct {
	server *Server
}

func (m *mutationResolver) createAccount(ctx context.Context, input AccountInput) (*Account, error) {
	return nil, nil
}

func (m *mutationResolver) createProduct(ctx context.Context, input ProductInput) (*Product, error) {
	return nil, nil
}

func (m *mutationResolver) createOrder(ctx context.Context, input OrderInput) (*Order, error) {
	return nil, nil
}
