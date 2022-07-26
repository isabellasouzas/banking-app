package account

import "banking-go/app/domain/types"

type Account struct {
	AccountID types.AccountID
	Name      string
	CPF       string
	Balance   int
}
