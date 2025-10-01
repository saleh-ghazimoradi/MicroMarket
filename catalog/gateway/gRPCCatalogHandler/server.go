package gRPCCatalogHandler

import (
	"context"
	"github.com/saleh-ghazimoradi/MicroMarket/catalog/domain"
	"github.com/saleh-ghazimoradi/MicroMarket/catalog/dto"
	"github.com/saleh-ghazimoradi/MicroMarket/catalog/gateway/proto"
	"github.com/saleh-ghazimoradi/MicroMarket/catalog/service"
	"google.golang.org/grpc"
	"net"
)

type GRPCHandler struct {
	catalogService service.CatalogService
	proto.UnimplementedCatalogServiceServer
}

func (g *GRPCHandler) CreateCatalog(ctx context.Context, req *proto.CreateCatalogRequest) (*proto.CreateCatalogResponse, error) {
	catalog, err := g.catalogService.CreateCatalog(ctx, &dto.Catalog{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
	})
	if err != nil {
		return nil, err
	}
	return &proto.CreateCatalogResponse{
		Catalog: &proto.Catalog{
			Id:          catalog.Id,
			Name:        catalog.Name,
			Description: catalog.Description,
			Price:       catalog.Price,
		},
	}, nil
}

func (g *GRPCHandler) GetCatalogById(ctx context.Context, req *proto.GetCatalogRequest) (*proto.GetCatalogResponse, error) {
	catalog, err := g.catalogService.GetCatalogById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &proto.GetCatalogResponse{
		Catalog: &proto.Catalog{
			Id:          catalog.Id,
			Name:        catalog.Name,
			Description: catalog.Description,
			Price:       catalog.Price,
		},
	}, nil
}

func (g *GRPCHandler) GetCatalogs(ctx context.Context, req *proto.GetCatalogsRequest) (*proto.GetCatalogsResponse, error) {
	var res []*domain.Catalog
	var err error
	if req.Query != "" {
		res, err = g.catalogService.SearchCatalog(ctx, req.Query, req.Offset, req.Limit)
	} else if len(req.Ids) != 0 {
		res, err = g.catalogService.GetCatalogsByIds(ctx, req.Ids)
	} else {
		res, err = g.catalogService.GetCatalogs(ctx, req.Offset, req.Limit)
	}
	if err != nil {
		return nil, err
	}
	catalogs := make([]*proto.Catalog, len(res))
	for _, c := range res {
		catalogs = append(catalogs, &proto.Catalog{
			Id:          c.Id,
			Name:        c.Name,
			Description: c.Description,
			Price:       c.Price,
		})
	}
	return &proto.GetCatalogsResponse{
		Catalogs: catalogs,
	}, nil
}

func (g *GRPCHandler) Serve(addr string) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()
	proto.RegisterCatalogServiceServer(grpcServer, g)
	return grpcServer.Serve(lis)
}

func NewGrpcHandler(catalogService service.CatalogService) *GRPCHandler {
	return &GRPCHandler{
		catalogService: catalogService,
	}
}
