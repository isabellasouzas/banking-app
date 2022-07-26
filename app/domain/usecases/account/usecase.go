package account

import (
	"context"

	"banking-app/app/domain/entities/account"
	"banking-app/app/domain/types"
)

type Usecase interface {
	Create(context.Context, account.Account) (account.Account, error)
	GetByAccountID(context.Context, types.AccountID) (account.Account, error)
}

type UseCase struct {
	accountRepo Repository
}

func NewUseCase(repo Repository) *UseCase {
	return &UseCase{
		accountRepo: repo,
	}
}
