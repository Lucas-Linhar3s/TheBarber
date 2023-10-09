package supabase

import (
	"encoding/json"
	"fmt"

	"github.com/Lucas-Linhar3s/TheBarber/database"
	"github.com/Lucas-Linhar3s/TheBarber/infrastructure/account"
)

type PGRepository struct {
	DB *database.Client
}

func (r *PGRepository) Login(model account.Account) (token string, err error) {
	data, _, err := r.DB.From(model.Email).Select("*", "exact", false).Execute()
	if err != nil {
		return "", err
	}

	var v *account.Account

	datas := json.Unmarshal(data, v)
	if datas != nil {
		return "", datas
	}

	return fmt.Sprintf("format", v), nil
}

func (r *PGRepository) CreateAccount(model *account.Account) (createdId string, err error) {

	return "", nil
}
