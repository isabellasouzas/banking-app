package transfer

import (
	"context"

	"github/isabellasouzas/banking-go/app/domain/entities/transfer"
	"github/isabellasouzas/banking-go/app/domain/types"
)

func (u usecase) GetByAccount(ctx context.Context, accountID types.AccountID) ([]transfer.Transfer, error) {

	return nil, nil
}
