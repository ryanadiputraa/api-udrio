package service

import (
	"context"

	"github.com/ryanadiputraa/api-udrio/domain"
	"github.com/ryanadiputraa/api-udrio/pkg/pagination"
)

type ProductService struct {
	productRepository domain.IProductRepository
}

func NewProductService(repository domain.IProductRepository) domain.IProductService {
	return &ProductService{productRepository: repository}
}

func (s *ProductService) GetProductList(ctx context.Context, size int, page int, category int) (products []domain.Product, meta pagination.Page, err error) {
	if size <= 0 {
		size = 20
	}
	if page <= 0 {
		page = 1
	}

	offset := pagination.Offset(size, page)
	products, count, err := s.productRepository.GetProductList(ctx, size, offset, category)
	if err != nil {
		return nil, pagination.Page{}, err
	}

	pages := *pagination.NewPagination(size, page, int(count))
	meta = pagination.Page{
		CurrentPage:  pages.Page,
		TotalPage:    pages.TotalPage,
		TotalData:    pages.TotalData,
		NextPage:     pages.NextPage(),
		PreviousPage: pages.PreviousPage(),
	}

	return products, meta, nil
}

func (s *ProductService) GetProductDetail(ctx context.Context, productID string) (domain.Product, error) {
	product, err := s.productRepository.GetProduct(ctx, productID)
	if err != nil {
		return product, err
	}
	return product, nil
}

func (s *ProductService) GetProductCategoryList(ctx context.Context) ([]domain.ProductCategory, error) {
	categories, err := s.productRepository.GetProductCategoryList(ctx)
	if err != nil {
		return nil, err
	}
	return categories, nil
}
