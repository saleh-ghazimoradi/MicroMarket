package main

import (
	"github.com/saleh-ghazimoradi/MicroMarket/account/config"
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

	accountRepository := repository.NewAccountRepository(db, db)
	accountService := service.NewAccountService(accountRepository)
}
