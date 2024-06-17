package service

import (
	"errors"
	"fmt"
	"os"

	"go-online-store/internal/domain/customer/model"
	"go-online-store/internal/domain/customer/repository"
	"go-online-store/pkg/logger"
)

type UserService struct {
	repoCustomer repository.UserRepositoryImpl
	logger       *logger.Logger
}

type CustomerServiceImpl interface {
	CustomerRegister(user model.Customer) (*model.Customer, error)
	CustomerLogin(email string, password string) (*model.Customer, error)
}

func NewInstanceUserService() CustomerServiceImpl {
	log := logger.NewLogger(os.Stdout, "Service [Customer] :")
	customerRepo, err := repository.NewInstanceUserRepo()
	if err != nil {
		log.Error("Failed to initialize customer repository: " + err.Error())
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
		userService.logger.Error("User with email already exists:" + user.Email)
		return nil, errors.New("user with this email already exists")
	}

	user, err = userService.repoCustomer.CreateUser(user)
	if err != nil {
		userService.logger.Error("Failed to create user:" + err.Error())
		return nil, err
	}

	userService.logger.Info("Customer registered successfully:" + user.Email)
	return &user, nil
}

// UserLogin authenticates a user based on email and password
func (userService *UserService) CustomerLogin(email string, password string) (*model.Customer, error) {
	userService.logger.Info("Logging in customer with email: " + email)
	user, err := userService.repoCustomer.GetUserByEmail(email)
	if err != nil {
		userService.logger.Error("Failed to find user")
		return nil, fmt.Errorf("failed to find user: %w", ErrUserNotFound)
	}

	// Check if the provided password matches the user's stored password hash
	err = user.CheckPassword(password)
	if err != nil {
		userService.logger.Error("Invalid password for user: " + email)
		return nil, fmt.Errorf("invalid password: %w", ErrInvalidPassword)
	}

	userService.logger.Info("Customer logged in successfully: " + email)
	return &user, nil
}
