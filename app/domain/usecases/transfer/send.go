package transfer

import (
	"context"

	"banking-app/app/domain/entities/transfer"
)

// SendMoney transferReqBody sends money to destination account from origin account
func (u usecase) SendMoney(context.Context, transfer.Transfer) (transfer.Info, error) {
	const operation = `Usecase.Transfer.SendMoney`

	//todo validate if account exist
	//todo repositorio
	//tr, err := u.repo.sendMoney(ctx, transfer)
	//if err != nil {
	//	return transfer.TransferInfo{}, err
	//}
	return transfer.Info{}, nil
}
