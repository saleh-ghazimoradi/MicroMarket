package service

import (
	"context"
	"github.com/saleh-ghazimoradi/MicroMarket/account/domain"
	"github.com/saleh-ghazimoradi/MicroMarket/account/dto"
	"github.com/saleh-ghazimoradi/MicroMarket/account/repository"
	"github.com/segmentio/ksuid"
)

type AccountService interface {
	CreateAccount(ctx context.Context, input *dto.Account) (*domain.Account, error)
	GetAccount(ctx context.Context, id string) (*domain.Account, error)
	GetAccounts(ctx context.Context, offset, limit uint64) ([]*domain.Account, error)
}

type accountService struct {
	accountRepository repository.AccountRepository
}

func (a *accountService) CreateAccount(ctx context.Context, input *dto.Account) (*domain.Account, error) {
	var account domain.Account
	if err := a.accountRepository.CreateAccount(
		ctx,
		&domain.Account{
			Id:   ksuid.New().String(),
			Name: input.Name,
		}); err != nil {
		return nil, err
	}
	return &account, nil
}

func (a *accountService) GetAccount(ctx context.Context, id string) (*domain.Account, error) {
	return a.accountRepository.GetAccountById(ctx, id)
}

func (a *accountService) GetAccounts(ctx context.Context, offset, limit uint64) ([]*domain.Account, error) {
	return a.accountRepository.GetAccounts(ctx, offset, limit)
}

func NewAccountService(accountRepository repository.AccountRepository) AccountService {
	return &accountService{
		accountRepository: accountRepository,
	}
}
