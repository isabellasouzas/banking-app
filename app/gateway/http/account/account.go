package account

import (
	"github/isabellasouzas/banking-go/app/domain/usecases/account"
)

type Handler struct {
	Usecase account.Usecase
}

func NewHandler(accUsecase account.Usecase) Handler {
	return Handler{
		Usecase: accUsecase,
	}
}

type Account struct {
	AccountID string `json:"account_id"`
	Name      string `json:"name"`
	CPF       string `json:"cpf"`
	Balance   int    `json:"balance"`
}
