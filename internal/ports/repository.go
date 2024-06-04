package ports

import "payment-system-two/internal/models"

type Repository interface {
	FindUserByEmail(email string) (*models.User, error)
	TokenInBlacklist(token *string) bool
	CreateUser(user *models.User) error
	UpdateUser(user *models.User) error
	CreateAdmin(admin *models.Admin) error
	FindAdminByEmail(email string) (*models.Admin, error)
	UpdateAdmin(admin *models.Admin) error
	FindUserByAccNo(accountNo int) (*models.User, error)
	Transferfunds(user *models.User, recipient *models.User, amount float64) error
	ADDfunds(user *models.User, amount float64) error
}
