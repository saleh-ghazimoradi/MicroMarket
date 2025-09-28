package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/saleh-ghazimoradi/MicroMarket/graphql/config"
	"log"
	"net/http"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	server, err := NewGraphqlServer(cfg.Application.AccountURL, cfg.Application.CatalogURL, cfg.Application.OrderURL)
	if err != nil {
		log.Fatalf("server initialization failed: %v", err)
	}

	http.Handle("/graphql", handler.New(server.ToExecutableSchema()))
	http.Handle("/playground", playground.Handler("GraphQL playground", "/graphql"))
	log.Println("connect to http://localhost:8080/graphql")
	log.Fatal(http.ListenAndServe(cfg.Server.Port, nil))

}
