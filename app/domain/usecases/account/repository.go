package account

import (
	"context"

	"github/isabellasouzas/banking-go/app/domain/entities/account"
	"github/isabellasouzas/banking-go/app/domain/types"
)

type Repository interface {
	CreateAccount(context.Context, account.Account) (account.Account, error)
	GetByAccountID(context.Context, types.AccountID) (account.Account, error)
	Update(context.Context, account.Account) (account.Account, error)
}
