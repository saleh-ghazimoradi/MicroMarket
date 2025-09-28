package repository

import (
	"context"
	"database/sql"
	"github.com/saleh-ghazimoradi/MicroMarket/account/domain"
)

type AccountRepository interface {
	CreateAccount(ctx context.Context, account *domain.Account) error
	GetAccountById(ctx context.Context, id string) (*domain.Account, error)
	GetAccounts(ctx context.Context, offset, limit uint64) ([]*domain.Account, error)
}

type accountRepository struct {
	dbRead  *sql.DB
	dbWrite *sql.DB
}

func (c *accountRepository) CreateAccount(ctx context.Context, account *domain.Account) error {
	return nil
}

func (c *accountRepository) GetAccountById(ctx context.Context, id string) (*domain.Account, error) {
	return nil, nil
}

func (c *accountRepository) GetAccounts(ctx context.Context, offset, limit uint64) ([]*domain.Account, error) {
	return nil, nil
}

func NewAccountRepository(dbRead, dbWrite *sql.DB) AccountRepository {
	return &accountRepository{
		dbRead:  dbRead,
		dbWrite: dbWrite,
	}
}
