package repository

import (
	mysql "go-online-store/config/database/my_sql_db"
	"go-online-store/internal/domain/order/model"

	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

type OrderRepositoryImpl interface {
	CreateOrder(order *model.Order) error
}

func NewInstanceOrderRepository() (OrderRepositoryImpl, error) {
	db, err := mysql.ConnectDatabase()
	if err != nil {
		return nil, err
	}

	return &OrderRepository{db: db}, nil
}

func (orderRepo *OrderRepository) CreateOrder(order *model.Order) error {
	return orderRepo.db.Create(order).Error
}
