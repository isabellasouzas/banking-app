package transfer

import (
	"context"

	"banking-app/app/domain/entities/transfer"
	"banking-app/app/domain/types"
)

func (r Repository) GetByTransferID(ctx context.Context, id types.TransferID) (transfer.Info, error) {
	const operation = `Repository.Transfer.GetByTransferID`
	return transfer.Info{}, nil
}
