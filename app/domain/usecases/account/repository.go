package account

import (
	"context"

	"banking-app/app/domain/entities/account"
	"banking-app/app/domain/types"
)

type Repository interface {
	CreateAccount(context.Context, account.Account) (account.Account, error)
	GetByAccountID(context.Context, types.AccountID) (account.Account, error)
	Update(context.Context, account.Account) (account.Account, error)
}
