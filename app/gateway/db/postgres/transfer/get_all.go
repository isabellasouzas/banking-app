package transfer

import (
	"context"

	"banking-app/app/domain/entities/transfer"
)

func (r Repository) GetAll(ctx context.Context) ([]transfer.Info, error) {
	const operation = `Repository.Transfer.GetAll`
	return []transfer.Info{}, nil
}
