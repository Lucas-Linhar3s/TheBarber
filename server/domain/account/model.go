package account

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Phone       string    `json:"phone"`
	DateOfBirth string    `json:"date_of_birth"`
	Photo       string    `json:"photo"`
	IsBarber    bool      `json:"is_barber"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
