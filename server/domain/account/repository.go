package account

import (
	"github.com/Lucas-Linhar3s/TheBarber/database"
	"github.com/Lucas-Linhar3s/TheBarber/infrastructure/account"
	repo "github.com/Lucas-Linhar3s/TheBarber/infrastructure/account/supabase"
)

type repository struct {
	repo *repo.PGRepository
}

func newRepository(client *database.Client) *repository {
	return &repository{
		repo: &repo.PGRepository{DB: client},
	}
}

func (r *repository) ConverteForRepo(model Account) (account.Account, error) {
	return account.Account{
		Email:    model.Email,
		Password: model.Password,
	}, nil
}

func (r *repository) Login(model account.Account) (token string, err error) {
	return r.repo.Login(model)
}

func (r *repository) CreateAccount(model *account.Account) (createdId string, err error) {
	return r.repo.CreateAccount(model)
}
