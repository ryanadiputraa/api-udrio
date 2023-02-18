package handler

import (
	"context"

	"github.com/ryanadiputraa/api-udrio/domain"
	"github.com/ryanadiputraa/api-udrio/pkg/pagination"
)

type productHandler struct {
	productRepository domain.IProductRepository
}

func NewProductHandler(repository domain.IProductRepository) domain.IProductHandler {
	return &productHandler{productRepository: repository}
}

func (h *productHandler) GetProductList(ctx context.Context, size int, page int, category int) (products []domain.Product, meta pagination.Page, err error) {
	if size <= 0 {
		size = 20
	}
	if page <= 0 {
		page = 1
	}

	offset := pagination.Offset(size, page)
	products, count, err := h.productRepository.Fetch(ctx, size, offset, category)
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

func (h *productHandler) GetProductDetail(ctx context.Context, productID string) (domain.Product, error) {
	product, err := h.productRepository.FindByID(ctx, productID)
	if err != nil {
		return product, err
	}
	return product, nil
}

func (h *productHandler) GetProductCategoryList(ctx context.Context) ([]domain.ProductCategory, error) {
	categories, err := h.productRepository.FetchCategory(ctx)
	if err != nil {
		return nil, err
	}
	return categories, nil
}
