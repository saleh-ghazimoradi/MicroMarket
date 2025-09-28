package service

import (
	"context"
	"github.com/saleh-ghazimoradi/MicroMarket/account/domain"
	"github.com/saleh-ghazimoradi/MicroMarket/account/repository"
)

type AccountService interface {
	CreateAccount(ctx context.Context, name string) error
	GetAccount(ctx context.Context, id string) (*domain.Account, error)
	GetAccounts(ctx context.Context, offset, limit uint64) ([]*domain.Account, error)
}

type accountService struct {
	accountRepository repository.AccountRepository
}

func (a *accountService) CreateAccount(ctx context.Context, name string) error {
	return nil
}

func (a *accountService) GetAccount(ctx context.Context, id string) (*domain.Account, error) {
	return nil, nil
}

func (a *accountService) GetAccounts(ctx context.Context, offset, limit uint64) ([]*domain.Account, error) {
	return nil, nil
}

func NewAccountService(accountRepository repository.AccountRepository) AccountService {
	return &accountService{
		accountRepository: accountRepository,
	}
}
