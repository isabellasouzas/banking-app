package account

import (
	"context"

	"github/isabellasouzas/banking-go/app/domain/entities/account"
)

func (r Repository) Update(ctx context.Context, acc account.Account) (account.Account, error) {
	return account.Account{}, nil
}
