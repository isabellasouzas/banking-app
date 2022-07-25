package transfer

import (
	"context"

	"github/isabellasouzas/banking-go/app/domain/entities/transfer"
	"github/isabellasouzas/banking-go/app/domain/types"
)

type Repository interface {
	GetByAccount(context.Context, types.AccountID) ([]transfer.Info, error)
	GetByTransferID(context.Context, types.TransferID) (transfer.Info, error)
	SendMoney(ctx context.Context, transfer transfer.Transfer) (transfer.Info, error)
}
