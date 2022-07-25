package account

import (
	"context"

	"github/isabellasouzas/banking-go/app/domain/entities/account"
	"github/isabellasouzas/banking-go/app/domain/types"
)

func (r Repository) GetByAccountID(ctx context.Context, accID types.AccountID) (account.Account, error) {
	return account.Account{}, nil
}
