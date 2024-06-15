package service

import (
	"context"
	repoCart "go-online-store/internal/domain/cart/repository"
	"go-online-store/internal/domain/order/model"
	"go-online-store/internal/domain/order/repository"
	repoProduct "go-online-store/internal/domain/product/repository"
	"go-online-store/internal/middleware/jwt"
	customErrors "go-online-store/pkg/errors"

	"time"
)

type OrderService struct {
	repoOrder   repository.OrderRepositoryImpl
	repoCart    repoCart.CartRepositoryImpl
	repoProduct repoProduct.ProductRepositoryImpl
}

type OrderServiceImpl interface {
	Checkout(ctx context.Context) (*model.Order, error)
}

func NewOrderService() OrderServiceImpl {
	orderRepo, err := repository.NewInstanceOrderRepository()
	if err != nil {
		return nil
	}

	cartRepo, err := repoCart.NewCartRepository()
	if err != nil {
		return nil
	}

	productRepo, err := repoProduct.NewProductRepository()
	if err != nil {
		return nil
	}

	return &OrderService{
		repoOrder:   orderRepo,
		repoCart:    cartRepo,
		repoProduct: productRepo,
	}
}

// func (svcOrder *OrderService) Checkout(customerID uint, items []model.OrderItem) (*model.Order, error) {
// 	var total float64
// 	for _, item := range items {
// 		total += float64(item.Quantity) * item.Price
// 	}

// 	order := &model.Order{
// 		CustomerID: customerID,
// 		Total:      total,
// 		OrderDate:  time.Now(),
// 		Items:      items,
// 	}

// 	if err := svcOrder.repoOrder.CreateOrder(order); err != nil {
// 		return nil, customErrors.ErrInternalServerError
// 	}

// 	return order, nil
// }

func (svcOrder *OrderService) Checkout(ctx context.Context) (*model.Order, error) {
	customerCtx, ok := jwt.FromUser(ctx)
	if !ok {
		return nil, customErrors.ErrCustomerIDNotFound
	}

	// Get the cart for the customer
	cart, err := svcOrder.repoCart.GetCartByCustomerID(customerCtx.ID)
	if err != nil {
		return nil, err
	}

	// Calculate the total price
	var total float64
	for _, item := range cart.Items {
		product, err := svcOrder.repoProduct.GetByID(item.ProductID)
		if err != nil {
			return nil, err
		}
		total += float64(item.Quantity) * product.Price
	}

	// Create the order
	order := &model.Order{
		CustomerID: customerCtx.ID,
		Total:      total,
		OrderDate:  time.Now(),
		Items:      []model.OrderItem{},
	}

	// Create the order items
	for _, item := range cart.Items {
		orderItem := model.OrderItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		}
		order.Items = append(order.Items, orderItem)
	}

	// Save the order to the database
	err = svcOrder.repoOrder.CreateOrder(order)
	if err != nil {
		return nil, err
	}

	// Clear the cart
	err = svcOrder.repoCart.ClearCart(customerCtx.ID)
	if err != nil {
		return nil, err
	}

	return order, nil
}
