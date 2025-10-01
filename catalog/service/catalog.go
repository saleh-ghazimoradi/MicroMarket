package service

import (
	"context"
	"github.com/saleh-ghazimoradi/MicroMarket/catalog/domain"
	"github.com/saleh-ghazimoradi/MicroMarket/catalog/dto"
	"github.com/saleh-ghazimoradi/MicroMarket/catalog/repository"
	"github.com/segmentio/ksuid"
)

type CatalogService interface {
	CreateCatalog(ctx context.Context, input *dto.Catalog) (*domain.Catalog, error)
	GetCatalogById(ctx context.Context, id string) (*domain.Catalog, error)
	GetCatalogs(ctx context.Context, offset, limit uint64) ([]*domain.Catalog, error)
	GetCatalogsByIds(ctx context.Context, ids []string) ([]*domain.Catalog, error)
	SearchCatalog(ctx context.Context, query string, offset, limit uint64) ([]*domain.Catalog, error)
}

type catalogService struct {
	catalogRepository repository.CatalogRepository
}

func (c *catalogService) CreateCatalog(ctx context.Context, input *dto.Catalog) (*domain.Catalog, error) {
	var catalog domain.Catalog
	if err := c.catalogRepository.CreateCatalog(ctx, &domain.Catalog{
		Id:          ksuid.New().String(),
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
	}); err != nil {
		return nil, err
	}
	return &catalog, nil
}

func (c *catalogService) GetCatalogById(ctx context.Context, id string) (*domain.Catalog, error) {
	return c.catalogRepository.GetCatalogById(ctx, id)
}

func (c *catalogService) GetCatalogs(ctx context.Context, offset, limit uint64) ([]*domain.Catalog, error) {
	if limit > 100 || (offset == 0 && limit == 0) {
		limit = 100
	}
	return c.catalogRepository.GetCatalogs(ctx, offset, limit)
}

func (c *catalogService) GetCatalogsByIds(ctx context.Context, ids []string) ([]*domain.Catalog, error) {
	return c.catalogRepository.GetCatalogsByIds(ctx, ids)
}

func (c *catalogService) SearchCatalog(ctx context.Context, query string, offset, limit uint64) ([]*domain.Catalog, error) {
	if limit > 100 || (offset == 0 && limit == 0) {
		limit = 100
	}
	return c.catalogRepository.SearchCatalog(ctx, query, offset, limit)
}

func NewCatalogService(catalogRepository repository.CatalogRepository) CatalogService {
	return &catalogService{
		catalogRepository: catalogRepository,
	}
}
