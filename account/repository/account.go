package repository

import (
	"context"
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
	"github.com/saleh-ghazimoradi/MicroMarket/account/domain"
)

type AccountRepository interface {
	CreateAccount(ctx context.Context, account *domain.Account) error
	GetAccountById(ctx context.Context, id string) (*domain.Account, error)
	GetAccounts(ctx context.Context, offset, limit uint64) ([]*domain.Account, error)
	Close() error
}

type accountRepository struct {
	dbRead  *sql.DB
	dbWrite *sql.DB
}

func (c *accountRepository) CreateAccount(ctx context.Context, account *domain.Account) error {
	_, err := c.dbWrite.ExecContext(ctx, "INSERT INTO account(id,name) VALUES ($1,$2)", account.Id, account.Name)
	return err
}

func (c *accountRepository) GetAccountById(ctx context.Context, id string) (*domain.Account, error) {
	var account domain.Account
	query := `SELECT * FROM account WHERE id = $1`

	err := c.dbRead.QueryRowContext(ctx, query, id).Scan(&account.Id, &account.Name)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.New("account not found")
		default:
			return nil, err
		}
	}
	return &account, nil
}

func (c *accountRepository) GetAccounts(ctx context.Context, offset, limit uint64) ([]*domain.Account, error) {
	query := `SELECT * FROM account ORDER BY id DESC OFFSET $1 LIMIT $2`

	rows, err := c.dbRead.QueryContext(ctx, query, offset, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var accounts []*domain.Account
	for rows.Next() {
		var account domain.Account
		err := rows.Scan(&account.Id, &account.Name)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, &account)
	}
	return accounts, nil
}

func (c *accountRepository) Close() error {
	if err := c.dbRead.Close(); err != nil {
		return err
	}
	if err := c.dbWrite.Close(); err != nil {
		return err
	}
	return nil
}

func NewAccountRepository(dbRead, dbWrite *sql.DB) AccountRepository {
	return &accountRepository{
		dbRead:  dbRead,
		dbWrite: dbWrite,
	}
}
