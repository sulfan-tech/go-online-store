package jwt

import (
	"context"
	"fmt"
	"go-online-store/internal/domain/customer/model"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type contextKey string

const (
	customerKey      contextKey = "customer"
	customerEmailKey contextKey = "customer_email"
	customerIdKey    contextKey = "customer_id"
)

// Customer represents the customer information stored in the context.
type Customer struct {
	ID      uint
	Email   string
	Address string
}

// WithCustomer stores the customer information in the context.
func WithCustomer(ctx context.Context, customer Customer) context.Context {
	return context.WithValue(ctx, customerKey, customer)
}

// FromCustomer retrieves the customer information from the context.
func FromCustomer(ctx context.Context) (Customer, bool) {
	customer, ok := ctx.Value(customerKey).(Customer)
	return customer, ok
}

// GenerateJWT generates a JWT for the provided customer with a custom expiration time.
func GenerateJWT(customer *model.Customer, secret string, expiration time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": customer,
		"exp": time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateJWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Validate the JWT token
		authToken := c.Request().Header.Get("Authorization")
		splitToken := strings.Split(authToken, "Bearer ")

		if len(splitToken) < 2 {
			return c.JSON(http.StatusUnauthorized, "Customer Unauthorized")
		}

		token := splitToken[1]
		t, err := validateToken(token)

		if err != nil {
			return c.JSON(http.StatusUnauthorized, "Customer Unauthorized: "+err.Error())
		}

		claims, ok := t.Claims.(jwt.MapClaims)
		if !ok {
			return c.JSON(http.StatusInternalServerError, "Failed to get customer information from token")
		}

		subClaim, ok := claims["sub"].(map[string]interface{})
		if !ok {
			return c.JSON(http.StatusInternalServerError, "Failed to get sub claim")
		}

		emailFromSubClaim, ok := subClaim["email"].(string)
		if !ok {
			return c.JSON(http.StatusInternalServerError, "Failed to get email from sub claim")
		}

		addressFromSubClaim, ok := subClaim["address"].(string)
		if !ok {
			return c.JSON(http.StatusInternalServerError, "Failed to get address from sub claim")
		}

		// Extract and convert id from subClaim to uint
		idFloat, ok := subClaim["id"].(float64)
		if !ok {
			return c.JSON(http.StatusInternalServerError, "Failed to get id from sub claim")
		}
		userIdFromSubClaim := uint(idFloat)

		// Create a customer object
		customer := Customer{
			ID:      userIdFromSubClaim,
			Email:   emailFromSubClaim,
			Address: addressFromSubClaim,
		}

		ctx := WithCustomer(c.Request().Context(), customer)
		c.SetRequest(c.Request().WithContext(ctx))

		return next(c)
	}
}

// validateToken validates the JWT token.
func validateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// Use the secret key from the environment variable for validation
		secret := os.Getenv("JWT_SECRET")
		return []byte(secret), nil
	})
}
