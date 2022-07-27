package account

import (
	"net/http"

	"banking-app/app/domain/usecases/account"
	"banking-app/app/gateway/http/rest"
)

type Handler struct {
	Usecase account.Usecase
}

func NewHandler(accUsecase account.Usecase) Handler {
	return Handler{
		Usecase: accUsecase,
	}
}

type AccHandlers interface {
	Create(r *http.Request) rest.Response
	GetBalance(r *http.Request) rest.Response
}

type Account struct {
	AccountID string `json:"account_id"`
	Name      string `json:"name"`
	CPF       string `json:"cpf"`
	Balance   int    `json:"balance"`
}
