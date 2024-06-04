package repository

import (
	"payment-system-one/internal/models"
	"time"
)

func (p *Postgres) FindUserByEmail(email string) (*models.User, error) {
	user := &models.User{}

	if err := p.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// function to find user by ACC no
func (p *Postgres) FindUserByAccNo(accountNo int) (*models.User, error) {
	user := &models.User{}
	if err := p.DB.Where("account_no = ?", accountNo).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// create a user in thye database
func (p *Postgres) CreateUser(user *models.User) error {
	if err := p.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (p *Postgres) UpdateUser(user *models.User) error {
	if err := p.DB.Save(user).Error; err != nil {
		return err
	}
	return nil
}

// function for create admin
func (p *Postgres) CreateAdmin(admin *models.Admin) error {
	if err := p.DB.Create(admin).Error; err != nil {
		return err
	}
	return nil
}

// Function for finding admin by email
func (p *Postgres) FindAdminByEmail(email string) (*models.Admin, error) {
	admin := &models.Admin{}

	if err := p.DB.Where("email = ?", email).First(&admin).Error; err != nil {
		return nil, err
	}
	return admin, nil
}

// function for update Admin

func (p *Postgres) UpdateAdmin(admin *models.Admin) error {
	if err := p.DB.Save(admin).Error; err != nil {
		return err
	}
	return nil
}

// Transfer Funds function

func (p *Postgres) Transferfunds(user *models.User, recipient *models.User, amount float64) error {

	tx := p.DB.Begin()
	// deduct the amount from the payer

	user.AccountBalance -= amount
	// add the amount to the recipients acc

	recipient.AccountBalance += amount

	// save the transaction for the one paying

	if err := tx.Save(user).Error; err != nil {
		tx.Rollback()
		return err
	}

	// save the transaction for the recipient

	if err := tx.Save(recipient).Error; err != nil {
		tx.Rollback()
		return err
	}
	// save the transaction in the transaction table
	transaction := &models.Transaction{
		PayerAccount:      user.AccountNo,
		RecipientsAccount: recipient.AccountNo,
		TransactionType:   "debit",
		TransactionAmount: amount,
		TransactionDate:   time.Now(),
	}
	if err := tx.Create(transaction).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// Add Funds code
func (p *Postgres) ADDfunds(user *models.User, amount float64) error {

	tx := p.DB.Begin()

	// add the money to the user acc
	user.AccountBalance += amount

	// update the acc balance of the user by saving it
	if err := tx.Save(user).Error; err != nil {
		tx.Rollback()
		return err
	}
	// save the transaction in the transaction table
	transaction := &models.Transaction{
		PayerAccount:      user.AccountNo,
		TransactionType:   "debit",
		TransactionAmount: amount,
		TransactionDate:   time.Now(),
	}
	if err := tx.Create(transaction).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// Transaction History

func (p *Postgres) Transaction(account_no int) ([]models.Transaction, error) {
	var transaction []models.Transaction
	if err := p.DB.Where("payer_account = ? OR recipients_account = ? ", account_no, account_no).Find(&transaction).Error; err != nil {
		return nil, err
	}
	return transaction, nil
}
