package transfer

import (
	"context"

	"github/isabellasouzas/banking-go/app/domain/entities/transfer"
	"github/isabellasouzas/banking-go/app/domain/types"
)

type Usecase interface {
	GetAllTransfers(context.Context) ([]transfer.Info, error)
	GetByAccount(context.Context, types.AccountID) ([]transfer.Info, error)
	GetByTransferID(context.Context, types.TransferID) (transfer.Info, error)
	Send(context.Context, transfer.Transfer) (transfer.Info, error)
}

type usecase struct {
	repo Repository
}

func NewUsecase(repo Repository) *usecase {
	return &usecase{
		repo: repo,
	}
}
