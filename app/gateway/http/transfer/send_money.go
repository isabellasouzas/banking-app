package transfer

import (
	"fmt"
	"net/http"

	"banking-app/app/domain/entities/transfer"
	"banking-app/app/gateway/http/rest"

	"github.com/pkg/errors"
)

func (h Handler) SendMoney(r *http.Request) rest.Response {
	const operation = `Handler.Transfer.SendMoney`

	ctx := r.Context()
	var reqBody transferReqBody

	input := transfer.Transfer{
		OriginID:      reqBody.OriginID,
		DestinationID: reqBody.DestinationID,
		Amount:        reqBody.Amount,
	}

	resp, err := h.Usecase.SendMoney(ctx, input)
	if err != nil {
		errors.Wrap(err, operation)
		fmt.Println("couldn't finalize the transfer")
		return rest.Response{}
	}

	fmt.Println("transfer successfully concluded")

	return rest.OK(resp)

}
