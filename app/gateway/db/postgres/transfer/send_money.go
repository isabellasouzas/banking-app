package transfer

import (
	"context"

	"banking-app/app/domain/entities/transfer"
)

func (r Repository) SendMoney(ctx context.Context, transfer transfer.Transfer) (transfer.Info, error) {
	const operation = `Repository.Transfer.SendMoney`
	return transfer.Info{}, nil
}
