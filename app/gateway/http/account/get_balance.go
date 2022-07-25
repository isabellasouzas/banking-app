package account

import (
	"fmt"
	"net/http"
	"strings"

	"github/isabellasouzas/banking-go/app/domain/types"
	"github/isabellasouzas/banking-go/gateway/http/rest"

	"github/isabellasouzas/pkg/mod/github.com/gorilla/mux@v1.8.0"
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
