package main

import (
	"github.com/saleh-ghazimoradi/MicroMarket/account/config"
	"github.com/saleh-ghazimoradi/MicroMarket/account/gateway/grpcAccountHandler"
	"github.com/saleh-ghazimoradi/MicroMarket/account/migrations"
	"github.com/saleh-ghazimoradi/MicroMarket/account/repository"
	"github.com/saleh-ghazimoradi/MicroMarket/account/service"
	"github.com/saleh-ghazimoradi/MicroMarket/account/utils"
	"log"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	postgresql := utils.NewPostgresql(
		utils.WithHost(cfg.Postgresql.Host),
		utils.WithPort(cfg.Postgresql.Port),
		utils.WithName(cfg.Postgresql.Name),
		utils.WithUser(cfg.Postgresql.User),
		utils.WithPassword(cfg.Postgresql.Password),
		utils.WithMaxOpenConn(cfg.Postgresql.MaxOpenConn),
		utils.WithMaxIdleConn(cfg.Postgresql.MaxIdleConn),
		utils.WithMaxIdleTime(cfg.Postgresql.MaxIdleTime),
		utils.WithSSLMode(cfg.Postgresql.SSLMode),
	)

	db, err := postgresql.Connect()
	if err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("error closing database connection: %v", err)
		}
	}()

	migrator, err := migrations.NewMigrator(db, postgresql.Name)
	if err != nil {
		log.Fatalf("error initializing migrator: %v", err)
	}

	if err := migrator.Up(); err != nil {
		log.Fatalf("error initializing migrator up: %v", err)
	}

	defer func() {
		if err = migrator.Close(); err != nil {
			log.Fatalf("error closing migrator up: %v", err)
		}
	}()

	accountRepository := repository.NewAccountRepository(db, db)
	accountService := service.NewAccountService(accountRepository)
	accountHandler := grpcAccountHandler.NewGRPCHandler(accountService)

	log.Println("Server is running on port", cfg.AccountServer.GRPCPort)
	if err = accountHandler.Serve(cfg.AccountServer.GRPCPort); err != nil {
		log.Fatalf("error serving grpc handler: %v", err)
	}
}
