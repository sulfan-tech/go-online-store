package model

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Customer struct {
	ID               uint      `gorm:"column:id" json:"id"`
	UserName         string    `gorm:"column:username" json:"username"`
	Email            string    `gorm:"unique;not null" json:"email"`
	Password         string    `gorm:"not null" json:"-"`
	FullName         string    `gorm:"column:full_name" json:"full_name"`
	Phone            string    `gorm:"column:phone" json:"phone"`
	Address          string    `gorm:"column:address" json:"address"`
	City             string    `gorm:"column:city" json:"city"`
	PostalCode       string    `gorm:"column:postal_code" json:"postal_code"`
	Country          string    `gorm:"column:country" json:"country"`
	DateOfBirth      time.Time `gorm:"column:date_of_birth" json:"date_of_birth"`
	RegistrationDate time.Time `json:"registration_date"`
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
