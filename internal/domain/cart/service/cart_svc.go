package service

import (
	"context"
	"errors"
	"go-online-store/internal/domain/cart/model"
	"go-online-store/internal/domain/cart/repository"
	"go-online-store/internal/middleware/jwt"
	customErrors "go-online-store/pkg/errors"

	"gorm.io/gorm"
)

type CartService struct {
	repoCart repository.CartRepositoryImpl
}

type CartServiceImpl interface {
	AddToCart(ctx context.Context, productID, quantity uint) error
	GetCartByCustomerID(ctx context.Context) (*model.Cart, error)
	RemoveFromCart(ctx context.Context, productID uint) error
}

func NewInstanceCartService() CartServiceImpl {
	cartRepo, err := repository.NewCartRepository()
	if err != nil {
		return nil
	}
	return &CartService{
		repoCart: cartRepo,
	}
}

// AddToCart adds a product to the customer's shopping cart.
func (cartService *CartService) AddToCart(ctx context.Context, productID, quantity uint) error {
	customerCtx, ok := jwt.FromUser(ctx)
	if !ok {
		return customErrors.ErrCustomerIDNotFound
	}

	cart, err := cartService.repoCart.GetCartByCustomerID(customerCtx.ID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return customErrors.ErrFailedToRetrieveCart
	}

	// If cart doesn't exist, create a new cart
	if cart == nil {
		newCart := &model.Cart{
			CustomerID: customerCtx.ID,
		}
		if err := cartService.repoCart.CreateCart(newCart); err != nil {
			return customErrors.ErrFailedToCreateCart
		}
		cart = newCart
	}

	// Add the product to the cart
	if err := cartService.repoCart.CreateCartItem(cart.CartID, productID, quantity); err != nil {
		return customErrors.ErrFailedToAddToCart
	}

	return nil
}

// GetCartByCustomerID retrieves the cart for a specific customer.
func (cartService *CartService) GetCartByCustomerID(ctx context.Context) (*model.Cart, error) {
	customerCtx, ok := jwt.FromUser(ctx)
	if !ok {
		return nil, customErrors.ErrCustomerIDNotFound
	}

	cart, err := cartService.repoCart.GetCartByCustomerID(customerCtx.ID)
	if err != nil {
		return nil, customErrors.ErrFailedToRetrieveCart
	}

	return cart, nil
}

// RemoveFromCart removes a product from the customer's shopping cart.
func (cartService *CartService) RemoveFromCart(ctx context.Context, productID uint) error {
	customerCtx, ok := jwt.FromUser(ctx)
	if !ok {
		return customErrors.ErrCustomerIDNotFound
	}

	cart, err := cartService.repoCart.GetCartByCustomerID(customerCtx.ID)
	if err != nil {
		return customErrors.ErrFailedToRetrieveCart
	}

	if err := cartService.repoCart.DeleteCartItem(cart.CartID, productID); err != nil {
		return customErrors.ErrFailedToRemoveFromCart
	}

	return nil
}
