package models

import "gorm.io/gorm"
import "time"

type User struct {
	gorm.Model
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Password     string `json:"password"`
	DateOfBirth  string `json:"date_of_birth"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	Address      string `json:"address"`
	LoginCounter int    `json:"login_counter"`
	IsLocked     bool   `json:"is_locked"`
}

//type UserProfile struct {
//	gorm.Model
//	ValidIdentity string `json:"valid_identity"`
//	PassPort string `json:"passport"`
//
//}

type Transaction struct {
	gorm.Model
	UserID          uint    `json:"user_id"`
	Amount          float64 `json:"amount"`
	Reference       string  `json:"reference"`
	TransactionType string  `json:"transaction_type"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Admin struct {
	gorm.Model
	First_name   string    `json:"first_name"`
	Last_name    string    `json:"last_name"`
	User_name    string    `json:"user_name"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	LoginCounter int       `json:"login_counter"`
	IsLocked     bool      `json:"is_locked"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type AdminRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
