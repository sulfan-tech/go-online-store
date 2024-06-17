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
	UpdateOrder(order *model.Order) error
	GetOrderById(id uint) (*model.Order, error)
	CreateTransaction(transaction *model.Transaction) error
	UpdateTransaction(transaction *model.Transaction) error
	GetTransactionByID(id uint) (*model.Transaction, error)
}

func NewInstanceOrderRepository() (OrderRepositoryImpl, error) {
	db, err := mysql.ConnectDatabase()
	if err != nil {
		return nil, err
	}

	return &OrderRepository{db: db}, nil
}

func (orderRepo *OrderRepository) GetOrderById(id uint) (*model.Order, error) {
	var order model.Order
	if err := orderRepo.db.First(&order, id).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (orderRepo *OrderRepository) CreateOrder(order *model.Order) error {
	return orderRepo.db.Create(order).Error
}

func (orderRepo *OrderRepository) UpdateOrder(order *model.Order) error {
	return orderRepo.db.Save(order).Error
}

func (orderRepo *OrderRepository) CreateTransaction(transaction *model.Transaction) error {
	return orderRepo.db.Create(transaction).Error
}

func (orderRepo *OrderRepository) UpdateTransaction(transaction *model.Transaction) error {
	return orderRepo.db.Save(transaction).Error
}

func (orderRepo *OrderRepository) GetTransactionByID(id uint) (*model.Transaction, error) {
	var transaction model.Transaction
	if err := orderRepo.db.First(&transaction, id).Error; err != nil {
		return nil, err
	}
	return &transaction, nil
}
