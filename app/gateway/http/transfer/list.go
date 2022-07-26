package transfer

import (
	"fmt"
	"net/http"

	"banking-app/app/domain/types"
	"banking-app/app/gateway/http/rest"

	"github.com/pkg/errors"
)

func (h Handler) List(r *http.Request) rest.Response {
	const operation = `Handler.Transfer.List`

	ctx := r.Context()
	var reqBody transferReqBody

	resp, err := h.Usecase.List(ctx, types.TransferID(reqBody.Info.ID))
	if err != nil {
		errors.Wrap(err, operation)
		// TODO add log
		fmt.Println("couldn't get the transfer")
		return rest.Response{}
	}

	fmt.Println("get transfer successfully")
	return rest.OK(resp)
}
