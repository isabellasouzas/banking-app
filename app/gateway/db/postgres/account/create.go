package account

import (
	"context"

	"banking-app/app/domain/entities/account"
)

func (r Repository) Create(ctx context.Context, acc account.Account) (account.Account, error) {
	const operation = `Repository.Account.Create`

	const statment string = `
		INSERT into
			accounts (document, secret, name, balance)
		values 
			($1, $2, $3, $4)
		returning 
			id, external_id, created_at, updated_at
	`
	return account.Account{}, nil
}
