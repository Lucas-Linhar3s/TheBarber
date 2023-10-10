package account

import (
	"github.com/Lucas-Linhar3s/TheBarber/database"
)

type Service struct {
	pg IAccount
}

// NewService creates a new account service
func GetService(pg IAccount) *Service {
	return &Service{
		pg: pg,
	}
}

// GetRepository Responsavel por pegar se conectar com os metodos no banco de dados
func GetRepository(client *database.Client) IAccount {
	return newRepository(client)
}

// CreateAccount creates a new account
func (s *Service) CreateAccount(model *Account) (createdId string, err error) {
	dados, err := s.pg.ConvertForRepo(*model)
	if err != nil {
		return "", err
	}
	return s.pg.CreateAccount(&dados)
}

// Login logs in an account
func (s *Service) Login(model Account) (token string, err error) {
	dados, err := s.pg.ConvertForRepo(model)
	if err != nil {
		return "", err
	}

	return s.pg.Login(dados)
}
