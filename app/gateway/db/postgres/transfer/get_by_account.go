package transfer

import (
	"context"

	"github/isabellasouzas/banking-go/app/domain/entities/transfer"
	"github/isabellasouzas/banking-go/app/domain/types"
)

func (r Repository) GetByAccount(ctx context.Context, accID types.AccountID) ([]transfer.Info, error) {
	return []transfer.Info{}, nil
}
