package account

import (
	"context"

	"banking-app/app/domain/entities/account"
	"banking-app/app/domain/types"

	"github.com/pkg/errors"
)

func (u UseCase) GetByAccountID(ctx context.Context, accID types.AccountID) (account.Account, error) {
	const operation = `Usecase.Account.GetByAccountID`

	acc, err := u.accountRepo.GetByAccountID(ctx, accID)
	if err != nil {
		errors.Wrap(err, operation)
		return account.Account{}, err
	}

	return acc, nil
}
