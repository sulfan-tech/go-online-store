package customer

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"go-online-store/internal/domain/customer/model"
	"go-online-store/internal/handlers/customer"
)

// MockUserService is a mocked implementation of UserServiceImpl for testing.
type MockUserService struct {
	mock.Mock
}

// UserLogin mocks the UserLogin method of UserServiceImpl.
func (m *MockUserService) CustomerLogin(email, password string) (*model.Customer, error) {
	args := m.Called(email, password)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Customer), nil
}

// UserRegister mocks the UserRegister method of UserServiceImpl.
func (m *MockUserService) CustomerRegister(user model.Customer) (*model.Customer, error) {
	args := m.Called(user)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Customer), nil
}

// TestUserLogin tests the UserLogin handler function.
func TestUserLogin(t *testing.T) {
	// Create an instance of the mock service
	// _ = godotenv.Load()
	mockService := new(MockUserService)
	handler := customer.NewCustomerHandler(mockService)

	// Mock data
	mockUser := &model.Customer{
		Email:    "test@example.com",
		Password: "password",
		UserName: "Test User",
	}

	// Mock the service method
	mockService.On("UserLogin", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(mockUser, nil)

	// Setup Echo context
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(`{"email": "test@example.com", "password": "password"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Perform the request
	err := handler.CustomerLogin(c)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	// Unmarshal response
	var response map[string]interface{}
	assert.NoError(t, json.NewDecoder(rec.Body).Decode(&response))

	// Check response fields
	assert.Contains(t, response, "token")
	assert.Contains(t, response, "data")
	assert.Equal(t, mockUser.Email, response["data"].(map[string]interface{})["email"])
	assert.Equal(t, mockUser.UserName, response["data"].(map[string]interface{})["name"])

	// Assert mock expectations
	mockService.AssertExpectations(t)
}

// TestUserRegister tests the UserRegister handler function.
func TestUserRegister(t *testing.T) {
	// Create an instance of the mock service
	mockService := new(MockUserService)
	handler := customer.NewCustomerHandler(mockService)

	// Mock data
	mockUser := &model.Customer{
		Email:    "newuser@example.com",
		Password: "password",
		UserName: "New User",
	}

	// Mock the service method
	mockService.On("UserRegister", mock.Anything).Return(mockUser, nil)

	// Setup Echo context
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/register", strings.NewReader(`{"name": "New User", "email": "newuser@example.com", "password": "password"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Perform the request
	err := handler.CustomerRegister(c)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)

	// Unmarshal response
	var response map[string]interface{}
	assert.NoError(t, json.NewDecoder(rec.Body).Decode(&response))

	// Check response fields
	assert.Contains(t, response, "data")
	assert.Equal(t, mockUser.Email, response["data"].(map[string]interface{})["email"])
	assert.Equal(t, mockUser.UserName, response["data"].(map[string]interface{})["name"])

	// Assert mock expectations
	mockService.AssertExpectations(t)
}
