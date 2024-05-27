package repository

import "payment-system-one/internal/models"

func (p *Postgres) FindUserByEmail(email string) (*models.User, error) {
	user := &models.User{}

	if err := p.DB.Where("email = ?", email).First(&user).Error; err != nil {
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
