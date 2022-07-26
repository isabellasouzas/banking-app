package transfer

import (
	"context"

	"banking-app/app/domain/entities/transfer"
	"banking-app/app/domain/types"
)

type Repository interface {
	GetByTransferID(context.Context, types.TransferID) (transfer.Info, error)
	GetAll(ctx context.Context) ([]transfer.Info, error)
	SendMoney(ctx context.Context, transfer transfer.Transfer) (transfer.Info, error)
}
