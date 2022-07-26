package account

import (
	"context"

	"banking-app/app/domain/entities/account"

	"github.com/pkg/errors"
)

func (u UseCase) Create(ctx context.Context, acc account.Account) (account.Account, error) {
	const operation = `Usecase.account.Create`

	NewAccount, err := u.repo.CreateAccount(ctx, acc)
	if err != nil {
		errors.Wrap(err, operation)
		return account.Account{}, err
	}

	return NewAccount, nil
}
