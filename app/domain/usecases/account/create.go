package account

import (
	"context"

	"github/isabellasouzas/banking-go/app/domain/entities/account"
)

func (u UseCase) Create(ctx context.Context, acc account.Account) (account.Account, error) {

	NewAccount, err := u.repo.CreateAccount(ctx, acc)
	if err != nil {
		return account.Account{}, err
	}

	return NewAccount, nil
}
