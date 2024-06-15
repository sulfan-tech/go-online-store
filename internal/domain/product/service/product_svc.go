package service

import (
	"context"
	"go-online-store/internal/domain/product/model"
	"go-online-store/internal/domain/product/repository"
)

type ProductService struct {
	repoProduct repository.ProductRepositoryImpl
}

type ProductServiceImpl interface {
	GetProductListByCategory(ctx context.Context, category string) ([]*model.Product, error)
}

func NewInstanceProductService() ProductServiceImpl {
	productRepo, err := repository.NewProductRepository()
	if err != nil {
		return nil
	}

	return &ProductService{
		repoProduct: productRepo,
	}
}

func (productService *ProductService) GetProductListByCategory(ctx context.Context, category string) ([]*model.Product, error) {
	productList, err := productService.repoProduct.GetProductsByCategory(category)
	if err != nil {
		return nil, err
	}

	return productList, nil
}

func (productService *ProductService) GetProductById(ctx context.Context, productId uint) (*model.Product, error) {
	product, err := productService.repoProduct.GetByID(productId)
	if err != nil {
		return nil, err
	}
	return product, nil
}
