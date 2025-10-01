package grpcAccountHandler

import (
	"context"
	"github.com/saleh-ghazimoradi/MicroMarket/account/gateway/proto"
	"github.com/saleh-ghazimoradi/MicroMarket/account/service"
	"google.golang.org/grpc"
	"net"
)

type GRPCHandler struct {
	accountService service.AccountService
	proto.UnimplementedAccountServiceServer
}

func (g *GRPCHandler) CreateAccount(ctx context.Context, r *proto.CreateAccountRequest) (*proto.CreateAccountResponse, error) {
	account, err := g.accountService.CreateAccount(ctx, r.Name)
	if err != nil {
		return nil, err
	}
	return &proto.CreateAccountResponse{
		Account: &proto.Account{
			Id:   account.Id,
			Name: account.Name,
		},
	}, nil
}

func (g *GRPCHandler) GetAccountById(ctx context.Context, r *proto.GetAccountRequest) (*proto.GetAccountResponse, error) {
	account, err := g.accountService.GetAccount(ctx, r.Id)
	if err != nil {
		return nil, err
	}
	return &proto.GetAccountResponse{
		Account: &proto.Account{
			Id:   account.Id,
			Name: account.Name,
		},
	}, nil
}

func (g *GRPCHandler) GetAccounts(ctx context.Context, r *proto.GetAccountsRequest) (*proto.GetAccountsResponse, error) {
	acc, err := g.accountService.GetAccounts(ctx, r.Offset, r.Limit)
	if err != nil {
		return nil, err
	}
	accounts := make([]*proto.Account, 0, len(acc))
	for _, ac := range acc {
		accounts = append(accounts, &proto.Account{
			Id:   ac.Id,
			Name: ac.Name,
		})
	}
	return &proto.GetAccountsResponse{Accounts: accounts}, nil
}

func (g *GRPCHandler) Serve(addr string) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()
	proto.RegisterAccountServiceServer(grpcServer, g)
	return grpcServer.Serve(lis)
}

func NewGRPCHandler(accountService service.AccountService) *GRPCHandler {
	return &GRPCHandler{accountService: accountService}
}
