package account

import (
	"fmt"
	"net/http"

	"banking-app/app/domain/entities/account"
	"banking-app/app/domain/types"
	"banking-app/app/gateway/http/rest"
)

func (h Handler) Create(r *http.Request) rest.Response {
	ctx := r.Context()
	var reqBody Account

	input := toRequestBody(reqBody)
	response, err := h.Usecase.Create(ctx, input)
	if err != nil {
		fmt.Println("couldn't create new account")
		return rest.Response{}
	}

	return rest.Created(response)
}

func toRequestBody(reqBody Account) account.Account {
	return account.Account{
		AccountID: types.AccountID(reqBody.AccountID),
		Name:      reqBody.Name,
		CPF:       reqBody.CPF,
		Balance:   reqBody.Balance,
	}
}
