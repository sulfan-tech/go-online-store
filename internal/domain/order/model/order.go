package model

import (
	"time"
)

type Order struct {
	ID              uint        `json:"id"`
	CustomerID      uint        `json:"customer_id"`
	OrderBy         string      `json:"order_by"`
	OrderNumber     string      `json:"order_number"`
	OrderDate       time.Time   `json:"order_date"`
	Total           float64     `json:"total"`
	ShippingFee     float64     `json:"shipping_fee"`
	Subtotal        float64     `json:"subtotal"`
	Tax             float64     `json:"tax"`
	Discount        float64     `json:"discount"`
	OrderStatus     string      `json:"order_status"`
	PaymentID       uint        `json:"payment_id"`
	PaymentDate     time.Time   `json:"payment_date"`
	PaymentStatus   string      `json:"payment_status"`
	ShippingAddress string      `json:"shipping_address"`
	BillingAddress  string      `json:"billing_address"`
	Currency        string      `json:"currency"`
	Items           []OrderItem `json:"items"`
	CreatedAt       time.Time   `json:"created_at"`
	UpdatedAt       time.Time   `json:"updated_at"`
}

type OrderItem struct {
	ID           uint    `json:"id"`
	OrderID      uint    `json:"order_id"`
	ProductID    uint    `json:"product_id"`
	ProductName  string  `json:"product_name"`
	ProductPrice float64 `json:"product_price"`
	Quantity     uint    `json:"quantity"`
	Subtotal     float64 `json:"subtotal"` // Harga total untuk item ini (ProductPrice * Quantity)
}

type Transaction struct {
	ID            string    `gorm:"primary_key" json:"id"`
	OrderID       uint      `json:"order_id"`
	PaymentStatus string    `json:"payment_status"`
	PaymentDate   time.Time `json:"payment_date"`
	Amount        float64   `json:"amount"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (Order) TableName() string {
	return "Order"
}
func (OrderItem) TableName() string {
	return "OrderItem"
}
func (Transaction) TableName() string {
	return "Transaction"
}
