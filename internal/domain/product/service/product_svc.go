package service

import (
	"context"
	"fmt"
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

	fmt.Println("MASUK SINI")
	test := ctx.Value("id")
	fmt.Println(test)

	return productList, nil
}
