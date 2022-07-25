package transfer

import (
	"context"

	"github/isabellasouzas/banking-go/app/domain/entities/transfer"
	"github/isabellasouzas/banking-go/app/domain/types"
)

func (u usecase) GetByTransferID(ctx context.Context, id types.TransferID) (transfer.Info, error) {

	return transfer.Info{}, nil
}
