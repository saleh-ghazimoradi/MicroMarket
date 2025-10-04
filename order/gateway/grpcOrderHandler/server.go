package grpcOrderHandler

import (
	"context"
	"github.com/saleh-ghazimoradi/MicroMarket/order/gateway/proto"
	"github.com/saleh-ghazimoradi/MicroMarket/order/service"
	"google.golang.org/grpc"
	"net"
)

type GrpcOrderHandler struct {
	orderService service.OrderService
	proto.UnimplementedOrderServiceServer
}

func (g *GrpcOrderHandler) CreateOrder(ctx context.Context, req *proto.CreateOrderRequest) (*proto.CreateOrderResponse, error) {
	return nil, nil
}

func (g *GrpcOrderHandler) GetOrdersForAccount(ctx context.Context, req *proto.GetOrdersForAccountRequest) (*proto.GetOrdersForAccountResponse, error) {
	return nil, nil
}

func (g *GrpcOrderHandler) Serve(addr string) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()
	proto.RegisterOrderServiceServer(grpcServer, g)
	return grpcServer.Serve(lis)
}

func NewGrpcOrderHandler(orderService service.OrderService) *GrpcOrderHandler {
	return &GrpcOrderHandler{
		orderService: orderService,
	}
}
