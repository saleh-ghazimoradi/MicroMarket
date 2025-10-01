package gRPCCatalogHandler

import (
	"context"
	"github.com/saleh-ghazimoradi/MicroMarket/catalog/domain"
	"github.com/saleh-ghazimoradi/MicroMarket/catalog/dto"
	"github.com/saleh-ghazimoradi/MicroMarket/catalog/gateway/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCCatalogClient struct {
	conn   *grpc.ClientConn
	client proto.CatalogServiceClient
}

func (c *GRPCCatalogClient) CreateCatalog(ctx context.Context, input *dto.Catalog) (*domain.Catalog, error) {
	catalog, err := c.client.CreateCatalog(ctx, &proto.CreateCatalogRequest{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
	})
	if err != nil {
		return nil, err
	}

	return &domain.Catalog{
		Id:          catalog.Catalog.Id,
		Name:        catalog.Catalog.Name,
		Description: catalog.Catalog.Description,
		Price:       catalog.Catalog.Price,
	}, nil
}

func (c *GRPCCatalogClient) GetCatalogById(ctx context.Context, id string) (*domain.Catalog, error) {
	catalog, err := c.client.GetCatalogById(ctx, &proto.GetCatalogRequest{Id: id})
	if err != nil {
		return nil, err
	}
	return &domain.Catalog{
		Id:          catalog.Catalog.Id,
		Name:        catalog.Catalog.Name,
		Description: catalog.Catalog.Description,
		Price:       catalog.Catalog.Price,
	}, nil
}

func (c *GRPCCatalogClient) GetCatalogs(ctx context.Context, offset, limit uint64, ids []string, query string) ([]*domain.Catalog, error) {
	catalogs, err := c.client.GetCatalogs(ctx, &proto.GetCatalogsRequest{
		Ids:    ids,
		Offset: offset,
		Limit:  limit,
		Query:  query,
	})

	if err != nil {
		return nil, err
	}

	cats := make([]*domain.Catalog, len(catalogs.Catalogs))
	for _, cat := range catalogs.Catalogs {
		cats = append(cats, &domain.Catalog{
			Id:          cat.Id,
			Name:        cat.Name,
			Description: cat.Description,
			Price:       cat.Price,
		})
	}
	return cats, nil
}

func (c *GRPCCatalogClient) Close() error {
	return c.conn.Close()
}

func NewGrpcCatalogClient(addr string) (*GRPCCatalogClient, error) {
	client, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	catalogServiceClient := proto.NewCatalogServiceClient(client)
	return &GRPCCatalogClient{
		client: catalogServiceClient,
	}, nil
}
