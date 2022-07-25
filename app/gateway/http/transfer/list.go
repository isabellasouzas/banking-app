package transfer

import (
	"fmt"
	"net/http"

	"github/isabellasouzas/banking-go/app/domain/types"
	"github/isabellasouzas/banking-go/gateway/http/rest"
)

func (h Handler) List(r *http.Request) rest.Response {
	ctx := r.Context()
	var reqBody transferReqBody

	resp, err := h.Usecase.GetByTransferID(ctx, types.TransferID(reqBody.Info.ID))
	if err != nil {
		fmt.Println("couldn't get the transfer")
		return rest.Response{}
	}

	fmt.Println("get transfer successfully")
	return rest.OK(resp)
}
