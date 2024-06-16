package service

import (
	"context"
	"fmt"
	cartModel "go-online-store/internal/domain/cart/model"
	repoCart "go-online-store/internal/domain/cart/repository"
	"go-online-store/internal/domain/order/model"
	repoOrder "go-online-store/internal/domain/order/repository"
	repoProduct "go-online-store/internal/domain/product/repository"
	"go-online-store/internal/middleware/jwt"
	"go-online-store/pkg/constant"
	customErrors "go-online-store/pkg/errors"

	"time"

	"github.com/google/uuid"
)

type OrderService struct {
	repoOrder   repoOrder.OrderRepositoryImpl
	repoCart    repoCart.CartRepositoryImpl
	repoProduct repoProduct.ProductRepositoryImpl
}

type OrderServiceImpl interface {
	Checkout(ctx context.Context) (*model.Order, error)
	UpdatePaymentStatus(ctx context.Context, orderID uint) error
}

func NewOrderService() (OrderServiceImpl, error) {
	orderRepo, err := repoOrder.NewInstanceOrderRepository()
	if err != nil {
		return nil, fmt.Errorf("failed to create order repository instance: %w", err)
	}

	cartRepo, err := repoCart.NewCartRepository()
	if err != nil {
		return nil, fmt.Errorf("failed to create cart repository instance: %w", err)
	}

	productRepo, err := repoProduct.NewProductRepository()
	if err != nil {
		return nil, fmt.Errorf("failed to create product repository instance: %w", err)
	}

	return &OrderService{
		repoOrder:   orderRepo,
		repoCart:    cartRepo,
		repoProduct: productRepo,
	}, nil
}

func (svcOrder *OrderService) Checkout(ctx context.Context) (*model.Order, error) {
	// Retrieve customer information from context
	customerCtx, ok := jwt.FromCustomer(ctx)
	if !ok {
		return nil, customErrors.ErrCustomerIDNotFound
	}

	// Retrieve the customer's cart
	cart, err := svcOrder.repoCart.GetCartByCustomerID(customerCtx.ID)
	if err != nil {
		return nil, err
	}

	// Check if the cart is empty
	if cart.Items == nil || len(cart.Items) == 0 {
		return nil, customErrors.ErrCartIsEmpty
	}

	// Calculate subtotal
	var subtotal float64
	for _, item := range cart.Items {
		product, err := svcOrder.repoProduct.GetByID(item.ProductID)
		if err != nil {
			return nil, err
		}
		subtotal += float64(item.Quantity) * product.Price
	}

	// Calculate shipping fee (optionally with discount)
	shippingFee := calculateShippingFee(cart.Items, true) // Example: Apply discount if true

	// Calculate total before tax and discount
	total := subtotal + shippingFee

	// Apply tax
	tax := applyTax(subtotal)
	total += tax

	// Apply discount
	discount := applyDiscount(subtotal)
	total -= discount

	// Create the order object
	order := &model.Order{
		CustomerID:      customerCtx.ID,
		OrderNumber:     generateOrderNumber(),
		OrderBy:         customerCtx.Email,
		OrderDate:       time.Now(),
		Total:           total,
		ShippingFee:     shippingFee,
		Subtotal:        subtotal,
		Tax:             tax,
		Discount:        discount,
		OrderStatus:     constant.ORDER_STATUS_PENDING,
		PaymentStatus:   constant.PAYMENT_STATUS_PENDING,
		PaymentDate:     time.Now(),
		ShippingAddress: customerCtx.Address,
		BillingAddress:  customerCtx.Address,
		Currency:        "IDR",
		Items:           make([]model.OrderItem, 0, len(cart.Items)),
	}

	// Populate order items
	for _, item := range cart.Items {
		product, err := svcOrder.repoProduct.GetByID(item.ProductID)
		if err != nil {
			return nil, err
		}

		orderItem := model.OrderItem{
			ProductID:    item.ProductID,
			Quantity:     item.Quantity,
			ProductName:  product.Name,
			ProductPrice: product.Price,
			Subtotal:     float64(item.Quantity) * product.Price,
		}
		order.Items = append(order.Items, orderItem)
	}

	// Create the order in the database
	err = svcOrder.repoOrder.CreateOrder(order)
	if err != nil {
		return nil, err
	}

	transaction := &model.Transaction{
		ID:            generatePaymentId(),
		OrderID:       order.ID,
		PaymentStatus: constant.PAYMENT_STATUS_PENDING,
		PaymentDate:   time.Now(),
		Amount:        total,
	}

	err = svcOrder.repoOrder.CreateTransaction(transaction)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (svcOrder *OrderService) UpdatePaymentStatus(ctx context.Context, orderID uint) error {
	// Retrieve the order and update its payment status
	order, err := svcOrder.repoOrder.GetOrderById(orderID)
	if err != nil {
		return fmt.Errorf("failed to retrieve order: %w", err)
	}

	//baypass
	order.PaymentStatus = constant.PAYMENT_STATUS_PAID
	order.PaymentDate = time.Now()
	order.OrderDate = time.Now()

	// Validate the status
	if order.PaymentStatus != constant.PAYMENT_STATUS_PAID && order.PaymentStatus != constant.PAYMENT_STATUS_PENDING && order.PaymentStatus != constant.PAYMENT_STATUS_FAILED {
		return fmt.Errorf("invalid payment status: %s", order.PaymentStatus)
	}

	err = svcOrder.repoOrder.UpdateOrder(order)
	if err != nil {
		return fmt.Errorf("failed to update order: %w", err)
	}

	transaction, err := svcOrder.repoOrder.GetTransactionByID(order.PaymentID)
	if err != nil {
		return fmt.Errorf("failed to retrieve transaction: %w", err)
	}

	//baypass
	transaction.PaymentStatus = constant.PAYMENT_STATUS_PAID
	transaction.PaymentDate = time.Now()

	err = svcOrder.repoOrder.UpdateTransaction(transaction)
	if err != nil {
		return fmt.Errorf("failed to update transaction: %w", err)
	}

	customerCtx, ok := jwt.FromCustomer(ctx)
	if !ok {
		return customErrors.ErrCustomerIDNotFound
	}

	cart, err := svcOrder.repoCart.GetCartByCustomerID(customerCtx.ID)
	if err != nil {
		return err
	}

	// Update stock on products and delete cart items
	for _, item := range cart.Items {
		product, err := svcOrder.repoProduct.GetByID(item.ProductID)
		if err != nil {
			return err
		}

		if product.Stok < item.Quantity {
			return customErrors.ErrProductStockNotAvailable
		}

		newStock := product.Stok - item.Quantity
		err = svcOrder.repoProduct.UpdateStock(product.ID, newStock)
		if err != nil {
			return err
		}

		err = svcOrder.repoCart.DeleteCartItem(cart.ID, product.ID)
		if err != nil {
			return err
		}
	}

	// Clear the customer's cart after successful checkout
	err = svcOrder.repoCart.ClearCart(customerCtx.ID)
	if err != nil {
		return err
	}

	return nil
}

// Function to calculate shipping fee based on business logic
func calculateShippingFee(cartItems []cartModel.CartItem, applyDiscount bool) float64 {
	// Example: Business logic to calculate shipping fee
	baseShippingFee := 10000.00
	if applyDiscount {
		baseShippingFee -= 2000 // Example discount
	}

	return baseShippingFee
}

// Function to apply tax based on business logic
func applyTax(subtotal float64) float64 {
	return 0.1 * subtotal // Example: 10% tax
}

// Function to apply discount based on business logic
func applyDiscount(subtotal float64) float64 {
	return 0.05 * subtotal // Example: 5% discount
}

func generateOrderNumber() string {
	uuid := uuid.New()
	return fmt.Sprintf("ORD-%s", uuid.String())
}
func generatePaymentId() string {
	uuid := uuid.New()
	return fmt.Sprintf("ORD-%s", uuid.String())
}
