package account

import (
	"context"

	"github/isabellasouzas/banking-go/app/domain/entities/account"
	"github/isabellasouzas/banking-go/app/domain/types"
)

func (u UseCase) GetByAccountID(ctx context.Context, accID types.AccountID) (account.Account, error) {

	acc, err := u.repo.GetByAccountID(ctx, accID)
	if err != nil {
		return account.Account{}, err
	}

	return acc, nil
}
