package customer

import (
	"errors"
	"net/http"
	"strings"
	"time"

	jwtConfig "go-online-store/config/jwt"
	"go-online-store/internal/domain/customer/model"
	"go-online-store/internal/domain/customer/service"
	"go-online-store/internal/middleware/jwt"

	"github.com/labstack/echo/v4"
)

type CustomerHandler struct {
	customerService service.CustomerServiceImpl
}

func NewCustomerHandler(customerService service.CustomerServiceImpl) *CustomerHandler {
	return &CustomerHandler{
		customerService: customerService,
	}
}

// CustomerLogin handles user authentication
func (h *CustomerHandler) CustomerLogin(c echo.Context) error {

	var loginRequest LoginRequest
	if err := c.Bind(&loginRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	// Call customerService to authenticate user
	customerAuth, err := h.customerService.CustomerLogin(loginRequest.Email, loginRequest.Password)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrUserNotFound):
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid email or password")
		case errors.Is(err, service.ErrInvalidPassword):
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid email or password")
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to authenticate user")
		}
	}

	// Generate JWT token for authenticated user
	token, err := jwt.GenerateJWT(customerAuth, jwtConfig.LoadJWTConfig().SecretKey, time.Hour*24)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to generate JWT")
	}

	response := map[string]interface{}{
		"token": token,
		"data":  customerAuth,
	}

	return c.JSON(http.StatusOK, response)

}

// CustomerRegister handles user registration
func (h *CustomerHandler) CustomerRegister(c echo.Context) error {
	var registerRequest RegisterRequest
	if err := c.Bind(&registerRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	// Create new user
	newUser := model.Customer{
		Email:    registerRequest.Email,
		Password: registerRequest.Password,
		UserName: registerRequest.Name,
	}

	newUser.SetPassword(registerRequest.Password)

	createdUser, err := h.customerService.CustomerRegister(newUser)
	if err != nil {
		// Handle specific error case where user with email already exists
		if strings.Contains(err.Error(), "user with this email already exists") {
			return echo.NewHTTPError(http.StatusConflict, "User with this email already exists")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create user")
	}

	response := map[string]interface{}{
		"data": createdUser,
	}

	return c.JSON(http.StatusCreated, response)
}
