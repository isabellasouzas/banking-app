package transfer

import (
	"context"

	"banking-app/app/domain/entities/transfer"
	"banking-app/app/domain/types"
)

type Usecase interface {
	List(context.Context, types.TransferID) (transfer.Info, error)
	SendMoney(context.Context, transfer.Transfer) (transfer.Info, error)
}

type usecase struct {
	repo Repository
}

func NewUsecase(repo Repository) *usecase {
	return &usecase{
		repo: repo,
	}
}
