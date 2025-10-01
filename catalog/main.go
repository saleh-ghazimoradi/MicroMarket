package main

import (
	"github.com/saleh-ghazimoradi/MicroMarket/catalog/config"
	"github.com/saleh-ghazimoradi/MicroMarket/catalog/gateway/gRPCCatalogHandler"
	"github.com/saleh-ghazimoradi/MicroMarket/catalog/repository"
	"github.com/saleh-ghazimoradi/MicroMarket/catalog/service"
	"github.com/saleh-ghazimoradi/MicroMarket/catalog/utils"
	"log"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	elastic := utils.NewElasticSearch(
		utils.WithHost(cfg.ElasticSearch.Host),
		utils.WithPort(cfg.ElasticSearch.Port),
		utils.WithUsername(cfg.ElasticSearch.Username),
		utils.WithPassword(cfg.ElasticSearch.Password),
		utils.WithTimeout(cfg.ElasticSearch.Timeout),
	)

	client, err := elastic.Connect()
	if err != nil {
		panic(err)
	}

	catalogRepository := repository.NewCatalogRepository(client, "catalogs")
	catalogService := service.NewCatalogService(catalogRepository)
	catalogHandler := gRPCCatalogHandler.NewGrpcHandler(catalogService)

	defer func() {
		if err = catalogRepository.Close(); err != nil {
			log.Fatalf("Error closing catalog repository: %v", err)
		}
	}()

	log.Println("Server is running on port", cfg.CatalogServer.Port)
	if err = catalogHandler.Serve(cfg.CatalogServer.Port); err != nil {
		log.Fatalf("error serving grpc handler: %v", err)
	}
}
