package account

import (
	"fmt"
	"net/http"
	"strings"

	"banking-app/app/domain/types"
	"banking-app/app/gateway/http/rest"

	"github.com/gorilla/mux"
)

func (h Handler) GetBalance(r *http.Request) rest.Response {
	ctx := r.Context()

	accID := types.AccountID(strings.TrimSpace(mux.Vars(r)["account-id"]))
	response, err := h.Usecase.GetByAccountID(ctx, accID)
	if err != nil {
		fmt.Println("Account ID is incorrect")
		return rest.Response{}
	}

	return rest.OK(response)
}
