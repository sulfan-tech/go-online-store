package repository

import (
	mysql "go-online-store/config/database/my_sql_db"
	"go-online-store/internal/domain/cart/model"

	"gorm.io/gorm"
)

type CartRepository struct {
	db *gorm.DB
}

type CartRepositoryImpl interface {
	GetCartByCustomerID(customerID uint) (*model.Cart, error)
	CreateCart(cart *model.Cart) error
	CreateCartItem(cartID, productID uint, quantity uint) error
	DeleteCartItem(cartID, productID uint) error
}

func NewCartRepository() (CartRepositoryImpl, error) {
	db, err := mysql.ConnectDatabase()
	if err != nil {
		return nil, err
	}

	return &CartRepository{db: db}, nil
}

// GetCartByCustomerID retrieves a cart by customer ID.
func (cartRepo *CartRepository) GetCartByCustomerID(customerID uint) (*model.Cart, error) {
	var cart model.Cart
	if err := cartRepo.db.Preload("Items").Where("customer_id = ?", customerID).First(&cart).Error; err != nil {
		return nil, err
	}
	return &cart, nil
}

// CreateCart initializes a new cart for a customer.
func (cartRepo *CartRepository) CreateCart(cart *model.Cart) error {
	return cartRepo.db.Create(cart).Error
}

// CreateCartItem adds a new item to the cart.
func (cartRepo *CartRepository) CreateCartItem(cartID, productID uint, quantity uint) error {
	cartItem := &model.CartItem{
		CartID:    cartID,
		ProductID: productID,
		Quantity:  quantity,
	}
	return cartRepo.db.Create(cartItem).Error
}

// DeleteCartItem removes an item from the cart.
func (cartRepo *CartRepository) DeleteCartItem(cartID, productID uint) error {
	return cartRepo.db.Delete(&model.CartItem{}, "cart_id = ? AND product_id = ?", cartID, productID).Error
}
