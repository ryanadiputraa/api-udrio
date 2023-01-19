package repository

import (
	"context"

	"github.com/ryanadiputraa/api-udrio/domain"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(conn *gorm.DB) domain.IProductRepository {
	return &ProductRepository{db: conn}
}

func (r *ProductRepository) GetProductList(ctx context.Context, page int, category int) ([]domain.Product, error) {
	var products []domain.Product

	modelQuery := r.db.Model(&domain.Product{}).Joins("ProductCategory", r.db.Where(&domain.ProductCategory{ID: category}))
	if category == 0 {
		modelQuery = r.db.Model(&domain.Product{}).Joins("ProductCategory")
	}

	err := modelQuery.Preload("ProductImages").Limit(50).Order("is_available desc, updated_at desc, created_at desc").Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}
