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

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) CustomerLogin(email, password string) (*model.Customer, error) {
	args := m.Called(email, password)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Customer), nil
}

func (m *MockUserService) CustomerRegister(user model.Customer) (*model.Customer, error) {
	args := m.Called(user)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Customer), nil
}

// TestUserLogin tests the UserLogin handler function.
func TestUserLogin(t *testing.T) {
	// _ = godotenv.Load()
	mockService := new(MockUserService)
	handler := customer.NewCustomerHandler(mockService)

	mockUser := &model.Customer{
		Email:    "test@example.com",
		Password: "password",
		UserName: "Test User",
	}

	mockService.On("UserLogin", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(mockUser, nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(`{"email": "test@example.com", "password": "password"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := handler.CustomerLogin(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var response map[string]interface{}
	assert.NoError(t, json.NewDecoder(rec.Body).Decode(&response))

	assert.Contains(t, response, "token")
	assert.Contains(t, response, "data")
	assert.Equal(t, mockUser.Email, response["data"].(map[string]interface{})["email"])
	assert.Equal(t, mockUser.UserName, response["data"].(map[string]interface{})["name"])

	mockService.AssertExpectations(t)
}

// TestUserRegister tests the UserRegister handler function.
func TestUserRegister(t *testing.T) {
	mockService := new(MockUserService)
	handler := customer.NewCustomerHandler(mockService)

	mockUser := &model.Customer{
		Email:    "newuser@example.com",
		Password: "password",
		UserName: "New User",
	}

	mockService.On("UserRegister", mock.Anything).Return(mockUser, nil)

	// Setup Echo context
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/register", strings.NewReader(`{"name": "New User", "email": "newuser@example.com", "password": "password"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := handler.CustomerRegister(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)

	var response map[string]interface{}
	assert.NoError(t, json.NewDecoder(rec.Body).Decode(&response))

	assert.Contains(t, response, "data")
	assert.Equal(t, mockUser.Email, response["data"].(map[string]interface{})["email"])
	assert.Equal(t, mockUser.UserName, response["data"].(map[string]interface{})["name"])

	mockService.AssertExpectations(t)
}
