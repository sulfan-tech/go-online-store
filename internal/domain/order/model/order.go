package model

import "time"

type Order struct {
	ID         uint        `json:"id"`
	CustomerID uint        `json:"customer_id"`
	Total      float64     `json:"total"`
	OrderDate  time.Time   `json:"order_date"`
	Items      []OrderItem `json:"items"`
}

type OrderItem struct {
	ID        uint `json:"id"`
	OrderID   uint `json:"order_id"`
	ProductID uint `json:"product_id"`
	// ProductName string `json:"product_name"`
	Quantity uint `json:"quantity"`
}

func (Order) TableName() string {
	return "Order"
}
func (OrderItem) TableName() string {
	return "OrderItem"
}
