package grpcOrderHandler

import (
	"github.com/saleh-ghazimoradi/MicroMarket/order/service"
	"google.golang.org/grpc"
	"net"
)

type GrpcOrderHandler struct {
	orderService service.OrderService
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
