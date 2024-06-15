package errors

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Custom error types
var (
	ErrBadRequest             = errors.New("bad request")
	ErrUnauthorized           = errors.New("unauthorized")
	ErrNotFound               = errors.New("not found")
	ErrInternalServerError    = errors.New("internal server error")
	ErrCustomerIDNotFound     = errors.New("customer ID not found")
	ErrInvalidCustomerID      = errors.New("invalid customer ID")
	ErrCartNotFound           = errors.New("cart not found")
	ErrProductAlreadyInCart   = errors.New("product already in cart")
	ErrFailedToCreateCart     = errors.New("failed to create cart")
	ErrFailedToAddToCart      = errors.New("failed to add to cart")
	ErrFailedToRemoveFromCart = errors.New("failed to remove from cart")
	ErrFailedToRetrieveCart   = errors.New("failed to retrieve cart")
)

// HTTPErrorHandler maps service errors to HTTP errors
func HTTPErrorHandler(err error) *echo.HTTPError {
	switch {
	case errors.Is(err, ErrBadRequest):
		return echo.NewHTTPError(http.StatusBadRequest, ErrBadRequest.Error())
	case errors.Is(err, ErrUnauthorized):
		return echo.NewHTTPError(http.StatusUnauthorized, ErrUnauthorized.Error())
	case errors.Is(err, ErrNotFound):
		return echo.NewHTTPError(http.StatusNotFound, ErrNotFound.Error())
	case errors.Is(err, ErrInternalServerError):
		return echo.NewHTTPError(http.StatusInternalServerError, ErrInternalServerError.Error())
	case errors.Is(err, ErrCustomerIDNotFound):
		return echo.NewHTTPError(http.StatusBadRequest, ErrCustomerIDNotFound.Error())
	case errors.Is(err, ErrInvalidCustomerID):
		return echo.NewHTTPError(http.StatusBadRequest, ErrInvalidCustomerID.Error())
	case errors.Is(err, ErrCartNotFound):
		return echo.NewHTTPError(http.StatusNotFound, ErrCartNotFound.Error())
	case errors.Is(err, ErrProductAlreadyInCart):
		return echo.NewHTTPError(http.StatusBadRequest, ErrProductAlreadyInCart.Error())
	case errors.Is(err, ErrFailedToCreateCart):
		return echo.NewHTTPError(http.StatusInternalServerError, ErrFailedToCreateCart.Error())
	case errors.Is(err, ErrFailedToAddToCart):
		return echo.NewHTTPError(http.StatusInternalServerError, ErrFailedToAddToCart.Error())
	case errors.Is(err, ErrFailedToRemoveFromCart):
		return echo.NewHTTPError(http.StatusInternalServerError, ErrFailedToRemoveFromCart.Error())
	case errors.Is(err, ErrFailedToRetrieveCart):
		return echo.NewHTTPError(http.StatusInternalServerError, ErrFailedToRetrieveCart.Error())
	default:
		return echo.NewHTTPError(http.StatusInternalServerError, ErrInternalServerError.Error())
	}
}
