package grpcAccountHandler

import (
	"context"
	"github.com/saleh-ghazimoradi/MicroMarket/account/domain"
	"github.com/saleh-ghazimoradi/MicroMarket/account/gateway/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCAccountClient struct {
	conn   *grpc.ClientConn
	client proto.AccountServiceClient
}

func (g *GRPCAccountClient) CreateAccount(ctx context.Context, name string) (*domain.Account, error) {
	resp, err := g.client.CreateAccount(ctx, &proto.CreateAccountRequest{Name: name})
	if err != nil {
		return nil, err
	}
	return &domain.Account{
		Id:   resp.Account.Id,
		Name: name,
	}, nil
}

func (g *GRPCAccountClient) GetAccountById(ctx context.Context, id string) (*domain.Account, error) {
	resp, err := g.client.GetAccountById(ctx, &proto.GetAccountRequest{Id: id})
	if err != nil {
		return nil, err
	}
	return &domain.Account{
		Id:   resp.Account.Id,
		Name: resp.Account.Name,
	}, nil
}

func (g *GRPCAccountClient) GetAccounts(ctx context.Context, offset, limit uint64) ([]*domain.Account, error) {
	resp, err := g.client.GetAccounts(ctx, &proto.GetAccountsRequest{Offset: offset, Limit: limit})
	if err != nil {
		return nil, err
	}
	accounts := make([]*domain.Account, len(resp.Accounts))
	for _, account := range resp.Accounts {
		accounts = append(accounts, &domain.Account{
			Id:   account.Id,
			Name: account.Name,
		})
	}
	return accounts, nil
}

func (g *GRPCAccountClient) Close() error {
	return g.conn.Close()
}

func NewGRPCAccountClient(addr string) (*GRPCAccountClient, error) {
	client, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	accountServiceClient := proto.NewAccountServiceClient(client)
	return &GRPCAccountClient{
		client: accountServiceClient,
	}, nil
}
