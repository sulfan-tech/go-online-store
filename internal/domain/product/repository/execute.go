package repository

import (
	mysql "go-online-store/config/database/my_sql_db"
	"go-online-store/internal/domain/product/model"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

type ProductRepositoryImpl interface {
	Create(product *model.Product) error
	Update(product *model.Product) error
	Delete(id uint) error
	GetByID(id uint) (*model.Product, error)
	GetProductsByCategory(category string) ([]*model.Product, error)
	GetAll() ([]*model.Product, error)
}

func NewProductRepository() (ProductRepositoryImpl, error) {
	db, err := mysql.ConnectDatabase()
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&model.Product{})
	return &ProductRepository{db}, nil
}

func (repo *ProductRepository) Create(product *model.Product) error {
	result := repo.db.Create(product)
	return result.Error
}

func (repo *ProductRepository) Update(product *model.Product) error {
	result := repo.db.Save(product)
	return result.Error
}

func (repo *ProductRepository) Delete(id uint) error {
	result := repo.db.Delete(&model.Product{}, id)
	return result.Error
}

func (repo *ProductRepository) GetByID(id uint) (*model.Product, error) {
	var product model.Product
	result := repo.db.First(&product, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

func (repo *ProductRepository) GetAll() ([]*model.Product, error) {
	var products []*model.Product
	result := repo.db.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}

func (repo *ProductRepository) GetProductsByCategory(category string) ([]*model.Product, error) {
	var products []*model.Product
	result := repo.db.Where("category = ?", category).Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}
