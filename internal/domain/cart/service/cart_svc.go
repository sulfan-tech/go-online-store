package service

import (
	"context"
	"errors"
	"go-online-store/internal/domain/cart/model"
	"go-online-store/internal/domain/cart/repository"
	repoProduct "go-online-store/internal/domain/product/repository"
	"go-online-store/internal/middleware/jwt"
	customErrors "go-online-store/pkg/errors"
	"go-online-store/pkg/logger"
	"os"
	"strconv"

	"gorm.io/gorm"
)

type CartService struct {
	repoCart    repository.CartRepositoryImpl
	repoProduct repoProduct.ProductRepositoryImpl
	logger      *logger.Logger
}

type CartServiceImpl interface {
	AddToCart(ctx context.Context, productID, quantity uint) error
	GetCartByCustomerID(ctx context.Context) (*model.Cart, error)
	RemoveFromCart(ctx context.Context, productID uint) error
}

func NewInstanceCartService() CartServiceImpl {
	log := logger.NewLogger(os.Stdout, "CartService")
	cartRepo, err := repository.NewCartRepository()
	if err != nil {
		log.Error("Failed to initialize cart repository: " + err.Error())
		return nil
	}

	productRepo, err := repoProduct.NewProductRepository()
	if err != nil {
		log.Error("Failed to initialize cart repository: " + err.Error())
		return nil
	}
	return &CartService{
		repoCart:    cartRepo,
		repoProduct: productRepo,
		logger:      log,
	}
}

// AddToCart adds a product to the customer's shopping cart.
func (cartService *CartService) AddToCart(ctx context.Context, productID, quantity uint) error {
	strPId := strconv.Itoa(int(productID))
	strQt := strconv.Itoa(int(quantity))

	cartService.logger.Info("Adding product to cart. ProductID:" + strPId + " Quantity:" + strQt)
	customerCtx, ok := jwt.FromCustomer(ctx)
	if !ok {
		cartService.logger.Error("CustomerID not found on ctx")
		return customErrors.ErrCustomerIDNotFound
	}

	cart, err := cartService.repoCart.GetCartByCustomerID(customerCtx.ID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		cartService.logger.Error("Failed to retrieve cart")
		return customErrors.ErrFailedToRetrieveCart
	}

	// If cart doesn't exist, create a new cart
	if cart == nil {
		newCart := &model.Cart{
			CustomerID: customerCtx.ID,
		}
		if err := cartService.repoCart.CreateCart(newCart); err != nil {
			cartService.logger.Error("Failed to create cart")
			return customErrors.ErrFailedToCreateCart
		}
		cart = newCart
	}

	// Check if product is available and add to cart
	product, err := cartService.repoProduct.GetByID(productID)
	if err != nil {
		cartService.logger.Error("Failed to retrieve product")
		return customErrors.ErrNotFound
	}

	if product.Stok < quantity {
		cartService.logger.Error("Product stock not available")
		return customErrors.ErrProductStockNotAvailable
	}

	// Add the product to the cart
	if err := cartService.repoCart.CreateCartItem(cart.ID, productID, quantity); err != nil {
		cartService.logger.Error("Failed to add product to cart")
		return customErrors.ErrFailedToAddToCart
	}

	return nil
}

// GetCartByCustomerID retrieves the cart for a specific customer.
func (cartService *CartService) GetCartByCustomerID(ctx context.Context) (*model.Cart, error) {
	cartService.logger.Info("Retrieving cart for customer")
	customerCtx, ok := jwt.FromCustomer(ctx)
	if !ok {
		cartService.logger.Error("CustomerID not found on ctx")
		return nil, customErrors.ErrCustomerIDNotFound
	}

	cart, err := cartService.repoCart.GetCartByCustomerID(customerCtx.ID)
	if err != nil {
		cartService.logger.Error("Failed to retrieve cart")
		return nil, customErrors.ErrFailedToRetrieveCart
	}

	cartService.logger.Info("Cart retrieved successfully")
	return cart, nil
}

// RemoveFromCart removes a product from the customer's shopping cart.
func (cartService *CartService) RemoveFromCart(ctx context.Context, productID uint) error {
	cartService.logger.Info("Removing product from cart")
	customerCtx, ok := jwt.FromCustomer(ctx)
	if !ok {
		cartService.logger.Error("CustomerID not found on ctx")
		return customErrors.ErrCustomerIDNotFound
	}

	cart, err := cartService.repoCart.GetCartByCustomerID(customerCtx.ID)
	if err != nil {
		cartService.logger.Error("Failed to retrieve cart")
		return customErrors.ErrFailedToRetrieveCart
	}

	if err := cartService.repoCart.DeleteCartItem(cart.ID, productID); err != nil {
		cartService.logger.Error("Failed to remove product from cart")
		return customErrors.ErrFailedToRemoveFromCart
	}

	cartService.logger.Info("Product removed from cart successfully")
	return nil
}
