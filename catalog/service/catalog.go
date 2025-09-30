package service

import (
	"github.com/saleh-ghazimoradi/MicroMarket/catalog/repository"
)

type CatalogService interface {
}

type catalogService struct {
	catalogRepository repository.CatalogRepository
}

func NewCatalogService(catalogRepository repository.CatalogRepository) CatalogService {
	return &catalogService{
		catalogRepository: catalogRepository,
	}
}
