package supabase

import (
	"github.com/Lucas-Linhar3s/TheBarber/database"
	"github.com/Lucas-Linhar3s/TheBarber/infrastructure/account"
)

type PGRepository struct {
	DB *database.Client
}

func (r *PGRepository) Login(model account.Account) (token string, err error) {
	return "", nil
}

func (r *PGRepository) CreateAccount(model *account.Account) (createdId string, err error) {
	_, count, err := r.DB.From("account").Insert(model, true, "", "id", "exact").Execute()
	if err != nil {
		return "", err
	}

	if count == 0 {
		return "", err
	}

	createdId = string(model.ID.String())
	return
}
