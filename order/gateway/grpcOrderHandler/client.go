package grpcOrderHandler

import (
	"context"
	"github.com/saleh-ghazimoradi/MicroMarket/order/domain"
	"github.com/saleh-ghazimoradi/MicroMarket/order/dto"
	"github.com/saleh-ghazimoradi/MicroMarket/order/gateway/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCOrderClient struct {
	conn   *grpc.ClientConn
	client proto.OrderServiceClient
}

func (g *GRPCOrderClient) CreateOrder(ctx context.Context, input *dto.Order) (*domain.Order, error) {
	return nil, nil
}

func (g *GRPCOrderClient) GetOrdersForAccount(ctx context.Context, accountId string) ([]*domain.Order, error) {
	return nil, nil
}

func (g *GRPCOrderClient) Close() error {
	return g.conn.Close()
}

func NewGRPCOrderClient(addr string) (*GRPCOrderClient, error) {
	client, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	orderServiceClient := proto.NewOrderServiceClient(client)
	return &GRPCOrderClient{
		client: orderServiceClient,
	}, nil
}
