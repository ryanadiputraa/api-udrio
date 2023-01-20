package service

import (
	"context"

	"github.com/ryanadiputraa/api-udrio/domain"
)

type ProductService struct {
	productRepository domain.IProductRepository
}

func NewProductService(repository domain.IProductRepository) domain.IProductService {
	return &ProductService{productRepository: repository}
}

func (s *ProductService) GetProductList(ctx context.Context, page int, category int) ([]domain.Product, error) {
	products, err := s.productRepository.GetProductList(ctx, page, category)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *ProductService) GetProductCategoryList(ctx context.Context) ([]domain.ProductCategory, error) {
	categories, err := s.productRepository.GetProductCategoryList(ctx)
	if err != nil {
		return nil, err
	}
	return categories, nil
}
