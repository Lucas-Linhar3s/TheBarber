package account

import "github.com/Lucas-Linhar3s/TheBarber/infrastructure/account"

type IAccount interface {
	CreateAccount(model *account.Account) (createdId string, err error)
	Login(model account.Account) (token string, err error)
	ConvertForRepo(model Account) (account.Account, error)
}
