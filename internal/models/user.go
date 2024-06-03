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
	AccountNo      int     `json:"accountNo"`
	AccountBalance float64 `json:"accountBalance"`
}

//type UserProfile struct {
//	gorm.Model
//	ValidIdentity string `json:"valid_identity"`
//	PassPort string `json:"passport"`
//
//}

type Transaction struct {
	gorm.Model
	PayerAccount      int       `json:"payerAccount"`
	RecipientsAccount int       `json:"recipientsAccount"`
	TransactionType   string    `json:"transactionType"`
	TransactionAmount float64   `json:"transactionAmount"`
	TransactionDate   time.Time `json:"transactionDate"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type AddMoney struct {
	Amount float64 `json:"amount"`
}

type TransferMoney struct {
	RecipiencACC int     `json:"recipienAcc"`
	Amount       float64 `json:"amount"`
}
