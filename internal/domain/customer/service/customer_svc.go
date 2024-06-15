package service

import (
	"errors"
	"fmt"

	"go-online-store/internal/domain/customer/model"
	"go-online-store/internal/domain/customer/repository"
)

type UserService struct {
	repoCustomer repository.UserRepositoryImpl
}

type CustomerServiceImpl interface {
	CustomerRegister(user model.Customer) (*model.Customer, error)
	CustomerLogin(email string, password string) (*model.Customer, error)
}

func NewInstanceUserService() CustomerServiceImpl {
	customerRepo, err := repository.NewInstanceUserRepo()
	if err != nil {
		return nil
	}
	return &UserService{
		repoCustomer: customerRepo,
	}
}

var (
	ErrUserNotFound    = errors.New("user not found")
	ErrInvalidPassword = errors.New("invalid password")
)

func (userService *UserService) CustomerRegister(user model.Customer) (*model.Customer, error) {
	// Check if user with the same email already exists
	_, err := userService.repoCustomer.GetUserByEmail(user.Email)
	if err == nil {
		return nil, errors.New("user with this email already exists")
	}

	// Create new user
	user, err = userService.repoCustomer.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// UserLogin authenticates a user based on email and password
func (userService *UserService) CustomerLogin(email string, password string) (*model.Customer, error) {
	// Retrieve user from repository by email
	user, err := userService.repoCustomer.GetUserByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("failed to find user: %w", ErrUserNotFound)
	}

	// Check if the provided password matches the user's stored password hash
	err = user.CheckPassword(password)
	if err != nil {
		return nil, fmt.Errorf("invalid password: %w", ErrInvalidPassword)
	}

	return &user, nil
}
