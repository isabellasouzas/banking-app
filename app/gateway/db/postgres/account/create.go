package account

import (
	"context"

	"banking-app/app/domain/entities/account"
)

func (r Repository) Create(ctx context.Context, acc account.Account) (account.Account, error) {
	const operation = `Repository.Account.Create`
	return account.Account{}, nil
}
