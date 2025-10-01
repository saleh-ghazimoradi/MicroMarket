package grpcOrderHandler

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCOrderClient struct {
	conn   *grpc.ClientConn
	client proto.OrderServiceClient
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
