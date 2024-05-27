package ports

import "payment-system-one/internal/models"

type Repository interface {
	FindUserByEmail(email string) (*models.User, error)
	TokenInBlacklist(token *string) bool
	CreateUser(user *models.User) error
	UpdateUser(user *models.User) error
	CreateAdmin(admin *models.Admin) error
	FindAdminByEmail(email string) (*models.Admin, error)
	UpdateAdmin(admin *models.Admin) error
}
