package model

import (
	"golang.org/x/crypto/bcrypt"
)

type Customer struct {
	ID       uint   `gorm:"column:customer_id" json:"customer_id"`
	UserName string `gorm:"column:username" json:"username"`
	Email    string `gorm:"unique;not null" json:"email"`
	Password string `gorm:"not null" json:"-"`
}

func (u *Customer) SetPassword(password string) error {
	hashedPassword, err := hashPassword(password)
	if err != nil {
		return err
	}
	u.Password = hashedPassword
	return nil
}

func (u *Customer) CheckPassword(password string) error {
	return comparePasswords(u.Password, password)
}

func (Customer) TableName() string {
	return "Customer"
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func comparePasswords(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
