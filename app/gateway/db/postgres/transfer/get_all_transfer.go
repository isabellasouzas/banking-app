package transfer

import (
	"context"

	"github/isabellasouzas/banking-go/app/domain/entities/transfer"
)

func (r Repository) GetAllTransfers(ctx context.Context) ([]transfer.Info, error) {
	return []transfer.Info{}, nil
}
