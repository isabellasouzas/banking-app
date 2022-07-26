package account

import (
	"context"

	"banking-app/app/domain/entities/account"
	"banking-app/app/domain/types"
)

func (r Repository) GetByAccountID(ctx context.Context, accID types.AccountID) (account.Account, error) {
	const operation = `Repository.Account.GetByAccountID`
	return account.Account{}, nil
}
