package account

import "github/isabellasouzas/banking-go/app/domain/types"

type Account struct {
	AccountID types.AccountID
	Name      string
	CPF       string
	Balance   int
}
