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

func (r *repository) ConvertForRepo(model Account) (account.Account, error) {
	return account.Account{
		ID:          model.ID,
		Name:        model.Name,
		LastName:    model.LastName,
		Phone:       model.Phone,
		DateOfBirth: model.DateOfBirth,
		Photo:       model.Photo,
		IsBarber:    model.IsBarber,
		Email:       model.Email,
		Password:    model.Password,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   *model.UpdatedAt,
	}, nil
}

func (r *repository) Login(model account.Account) (token string, err error) {
	return r.repo.Login(model)
}

func (r *repository) CreateAccount(model *account.Account) (createdId string, err error) {
	return r.repo.CreateAccount(model)
}
