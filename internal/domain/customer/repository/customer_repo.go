package repository

import (
	mysql "go-online-store/config/database/my_sql_db"
	"go-online-store/internal/domain/customer/model"
	"time"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

type UserRepositoryImpl interface {
	GetUserByEmail(email string) (model.Customer, error)
	CreateUser(user model.Customer) (model.Customer, error)
}

func NewInstanceUserRepo() (UserRepositoryImpl, error) {
	db, err := mysql.ConnectDatabase()
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&model.Customer{})
	return &UserRepository{db}, nil
}

func (customerSql *UserRepository) GetUserByEmail(email string) (model.Customer, error) {
	var user model.Customer
	err := customerSql.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (customerSql *UserRepository) CreateUser(customer model.Customer) (model.Customer, error) {
	customer.RegistrationDate = time.Now()
	err := customerSql.db.Create(&customer).Error
	if err != nil {
		return customer, err
	}
	return customer, nil
}
