package transfer

import (
	"fmt"
	"net/http"

	"github/isabellasouzas/banking-go/app/domain/entities/transfer"
	"github/isabellasouzas/banking-go/gateway/http/rest"
)

func (h Handler) SendMoney(r *http.Request) rest.Response {
	ctx := r.Context()
	var reqBody transferReqBody

	input := transfer.Transfer{
		OriginID:      reqBody.OriginID,
		DestinationID: reqBody.DestinationID,
		Amount:        reqBody.Amount,
	}

	resp, err := h.Usecase.Send(ctx, input)
	if err != nil {
		fmt.Println("couldn't finalize the transfer")
		return rest.Response{}
	}

	fmt.Println("transfer successfully concluded")

	return rest.OK(resp)

}
