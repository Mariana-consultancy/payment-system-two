package models

import "gorm.io/gorm"
import "time"

type User struct {
	gorm.Model
	FirstName      string  `json:"first_name"`
	LastName       string  `json:"last_name"`
	Password       string  `json:"password"`
	DateOfBirth    string  `json:"date_of_birth"`
	Email          string  `json:"email"`
	Phone          string  `json:"phone"`
	Address        string  `json:"address"`
	AccountNo      int     `json:"account_no"`
	AccountBalance float64 `json:"account_balance"`
}

//type UserProfile struct {
//	gorm.Model
//	ValidIdentity string `json:"valid_identity"`
//	PassPort string `json:"passport"`
//
//}

type Transaction struct {
	gorm.Model
	PayerAccount      int       `json:"payer_account"`
	RecipientsAccount int       `json:"recipients_account"`
	TransactionType   string    `json:"transaction_type"`
	TransactionAmount float64   `json:"transaction_amount"`
	TransactionDate   time.Time `json:"transaction_date"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type AddMoney struct {
	Amount float64 `json:"amount"`
}

type TransferMoney struct {
	AccountNo int     `json:"account_no"`
	Amount    float64 `json:"amount"`
}
