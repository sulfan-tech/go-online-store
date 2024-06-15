package cart

import (
	"context"
	"go-online-store/internal/domain/cart/service"
	customErrors "go-online-store/pkg/errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CartHandler struct {
	cartService service.CartServiceImpl
}

func NewCartHandler(cartService service.CartServiceImpl) *CartHandler {
	return &CartHandler{
		cartService: cartService,
	}
}

// AddToCartHandler handles the request to add a product to the cart
func (h *CartHandler) AddToCartHandler(c echo.Context) error {
	ctx := c.Request().Context()

	var req RequestAddToCard
	if err := c.Bind(&req); err != nil {
		return customErrors.HTTPErrorHandler(customErrors.ErrBadRequest)
	}

	if err := h.cartService.AddToCart(ctx, req.ProductID, req.Quantity); err != nil {
		return customErrors.HTTPErrorHandler(err)
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "product added to cart"})
}

// GetCartHandler handles the request to get the cart for a customer
func (h *CartHandler) GetCartHandler(c echo.Context) error {
	ctx := context.WithValue(c.Request().Context(), ctxKeyUserID, c.Get("id"))

	cart, err := h.cartService.GetCartByCustomerID(ctx)
	if err != nil {
		return customErrors.HTTPErrorHandler(err)
	}

	return c.JSON(http.StatusOK, cart)
}

// RemoveFromCartHandler handles the request to remove a product from the cart
func (h *CartHandler) RemoveFromCartHandler(c echo.Context) error {

	var req ReqRemoveCart
	if err := c.Bind(&req); err != nil {
		return customErrors.HTTPErrorHandler(customErrors.ErrBadRequest)
	}

	ctx := context.WithValue(c.Request().Context(), ctxKeyUserID, c.Get("id"))

	if err := h.cartService.RemoveFromCart(ctx, req.ProductID); err != nil {
		return customErrors.HTTPErrorHandler(err)
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "product removed from cart"})
}
