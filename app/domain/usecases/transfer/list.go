package transfer

import (
	"context"

	"banking-app/app/domain/entities/transfer"
	"banking-app/app/domain/types"
)

func (u usecase) List(context.Context, types.TransferID) (transfer.Info, error) {
	const operation = `Usecase.Transfer.List`
	return transfer.Info{}, nil
}
