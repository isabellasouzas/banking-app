package account

import (
	"context"

	"banking-app/app/domain/entities/account"
)

func (r Repository) Update(ctx context.Context, acc account.Account) error {
	const operation = `Repository.Account.Update`
	return nil
}
