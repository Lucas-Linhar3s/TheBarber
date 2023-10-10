package account

type Req struct {
	Name        string `json:"name" binding:"required"`
	LastName    string `json:"last_name" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Password    string `json:"password" binding:"required"`
	Phone       string `json:"phone" binding:"required"`
	DateOfBirth string `json:"date_of_birth" binding:"required"`
	Photo       string `json:"photo"`
	IsBarber    bool   `json:"is_barber"`
}
