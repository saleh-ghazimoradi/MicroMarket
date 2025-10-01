package repository

import (
	"context"
	"encoding/json"
	"github.com/saleh-ghazimoradi/MicroMarket/catalog/customErr"
	"github.com/saleh-ghazimoradi/MicroMarket/catalog/domain"
	"github.com/saleh-ghazimoradi/MicroMarket/catalog/dto"
	elastic "gopkg.in/olivere/elastic.v5"
)

type CatalogRepository interface {
	CreateCatalog(ctx context.Context, catalog *domain.Catalog) error
	GetCatalogById(ctx context.Context, id string) (*domain.Catalog, error)
	GetCatalogs(ctx context.Context, offset, limit uint64) ([]*domain.Catalog, error)
	GetCatalogsByIds(ctx context.Context, ids []string) ([]*domain.Catalog, error)
	SearchCatalog(ctx context.Context, query string, offset, limit uint64) ([]*domain.Catalog, error)
	Close() error
}

type catalogRepository struct {
	client *elastic.Client
	index  string
}

func (c *catalogRepository) CreateCatalog(ctx context.Context, catalog *domain.Catalog) error {
	_, err := c.client.Index().Index(c.index).Type("catalog").Id(catalog.Id).BodyJson(dto.Catalog{
		Name:        catalog.Name,
		Description: catalog.Description,
		Price:       catalog.Price,
	}).Do(ctx)
	return err
}

func (c *catalogRepository) GetCatalogById(ctx context.Context, id string) (*domain.Catalog, error) {
	res, err := c.client.Get().Index(c.index).Type("catalog").Id(id).Do(ctx)
	if err != nil {
		return nil, err
	}

	if !res.Found {
		return nil, customErr.ErrNotFound
	}

	var catalog dto.Catalog
	if err := json.Unmarshal(*res.Source, &catalog); err != nil {
		return nil, err
	}
	return &domain.Catalog{
		Id:          id,
		Name:        catalog.Name,
		Description: catalog.Description,
		Price:       catalog.Price,
	}, nil
}

func (c *catalogRepository) GetCatalogs(ctx context.Context, offset, limit uint64) ([]*domain.Catalog, error) {
	res, err := c.client.Search().Index(c.index).Type("catalog").Query(elastic.NewMatchAllQuery()).From(int(offset)).Size(int(limit)).Do(ctx)
	if err != nil {
		return nil, err
	}
	var catalogs []*domain.Catalog
	for _, hit := range res.Hits.Hits {
		var catalog dto.Catalog
		if err = json.Unmarshal(*hit.Source, &catalog); err != nil {
			return nil, err
		}
		catalogs = append(catalogs, &domain.Catalog{
			Id:          hit.Id,
			Name:        catalog.Name,
			Description: catalog.Description,
			Price:       catalog.Price,
		})
	}
	return catalogs, nil
}

func (c *catalogRepository) GetCatalogsByIds(ctx context.Context, ids []string) ([]*domain.Catalog, error) {
	var items []*elastic.MultiGetItem
	for _, id := range ids {
		items = append(items, elastic.NewMultiGetItem().
			Index(c.index).
			Type("catalog").
			Id(id),
		)
	}
	res, err := c.client.MultiGet().Add(items...).Do(ctx)
	if err != nil {
		return nil, err
	}
	var catalogs []*domain.Catalog
	for _, doc := range res.Docs {
		var catalog dto.Catalog
		if err = json.Unmarshal(*doc.Source, &catalog); err != nil {
			return nil, err
		}
		catalogs = append(catalogs, &domain.Catalog{
			Id:          doc.Id,
			Name:        catalog.Name,
			Description: catalog.Description,
			Price:       catalog.Price,
		})
	}
	return catalogs, nil
}

func (c *catalogRepository) SearchCatalog(ctx context.Context, query string, offset, limit uint64) ([]*domain.Catalog, error) {
	res, err := c.client.Search().Index(c.index).Type("catalog").Query(elastic.NewMultiMatchQuery(query, "name", "description")).From(int(offset)).Size(int(limit)).Do(ctx)
	if err != nil {
		return nil, err
	}
	var catalogs []*domain.Catalog
	for _, hit := range res.Hits.Hits {
		var catalog dto.Catalog
		if err = json.Unmarshal(*hit.Source, &catalog); err != nil {
			return nil, err
		}
		catalogs = append(catalogs, &domain.Catalog{
			Id:          hit.Id,
			Name:        catalog.Name,
			Description: catalog.Description,
			Price:       catalog.Price,
		})
	}
	return catalogs, nil
}

func (c *catalogRepository) Close() error {
	return nil
}

func NewCatalogRepository(client *elastic.Client, index string) CatalogRepository {
	return &catalogRepository{
		client: client,
		index:  index,
	}
}
