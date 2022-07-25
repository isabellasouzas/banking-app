package account

import (
	"context"

	"github/isabellasouzas/banking-go/app/domain/entities/account"
	"github/isabellasouzas/banking-go/app/domain/types"
)

type Usecase interface {
	Create(context.Context, account.Account) (account.Account, error)
	GetByAccountID(context.Context, types.AccountID) (account.Account, error)
}

type UseCase struct {
	repo Repository
}

func NewUseCase(repo Repository) *UseCase {
	return &UseCase{
		repo: repo,
	}
}
