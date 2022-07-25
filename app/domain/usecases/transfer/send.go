package transfer

import (
	"context"

	"github/isabellasouzas/banking-go/app/domain/entities/transfer"
)

// SendMoney transferReqBody sends money to destination account from origin account
func (u usecase) SendMoney(ctx context.Context, transfer transfer.Transfer) (transfer.Info, error) {

	//todo validate if account exist
	//todo repositorio
	//tr, err := u.repo.sendMoney(ctx, transfer)
	//if err != nil {
	//	return transfer.TransferInfo{}, err
	//}
	return transfer.Info{}, nil
}
