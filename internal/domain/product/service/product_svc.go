package service

import (
	"context"
	"fmt"
	"go-online-store/internal/domain/product/model"
	"go-online-store/internal/domain/product/repository"
	"go-online-store/pkg/logger"
	"os"
)

type ProductService struct {
	repoProduct repository.ProductRepositoryImpl
	logger      *logger.Logger
}

type ProductServiceImpl interface {
	GetProductListByCategory(ctx context.Context, category string) ([]*model.Product, error)
	GetProductById(ctx context.Context, productId uint) (*model.Product, error)
}

func NewInstanceProductService() ProductServiceImpl {
	log := logger.NewLogger(os.Stdout, "Service [Product] :")
	productRepo, err := repository.NewProductRepository()
	if err != nil {
		log.Error("Failed to initialize product repository: " + err.Error())
		return nil
	}

	return &ProductService{
		repoProduct: productRepo,
		logger:      log,
	}
}

func (productService *ProductService) GetProductListByCategory(ctx context.Context, category string) ([]*model.Product, error) {
	productService.logger.Info("Fetching products for category: " + category)
	productList, err := productService.repoProduct.GetProductsByCategory(category)
	if err != nil {
		productService.logger.Error("Failed to fetch products for category " + category + ": " + err.Error())
		return nil, err
	}

	return productList, nil
}

func (productService *ProductService) GetProductById(ctx context.Context, productId uint) (*model.Product, error) {
	productService.logger.Info("Fetching product with ID: " + fmt.Sprint(productId))
	product, err := productService.repoProduct.GetByID(productId)
	if err != nil {
		productService.logger.Error("Failed to fetch product with ID " + fmt.Sprint(productId) + ": " + err.Error())
		return nil, err
	}
	return product, nil
}
