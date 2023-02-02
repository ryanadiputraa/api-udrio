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

func (r *ProductRepository) GetProductList(ctx context.Context, size int, offset int, category int) (products []domain.Product, count int64, err error) {
	r.db.Model(&domain.Product{}).Count(&count)

	modelQuery := r.db.Model(&domain.Product{}).Joins("ProductCategory", r.db.Where(&domain.ProductCategory{ID: category}))
	if category == 0 {
		modelQuery = r.db.Model(&domain.Product{}).Joins("ProductCategory")
	}

	err = modelQuery.Preload("ProductImages").Limit(size).Offset(offset).Order("is_available desc, updated_at desc, created_at desc").Find(&products).Error
	if err != nil {
		return nil, 0, err
	}

	return products, count, nil
}

func (r *ProductRepository) GetProduct(ctx context.Context, productID string) (domain.Product, error) {
	var product domain.Product

	err := r.db.Model(&domain.Product{}).Joins("ProductCategory").Where(&domain.Product{ID: productID}).Preload("ProductImages").First(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *ProductRepository) GetProductCategoryList(ctx context.Context) ([]domain.ProductCategory, error) {
	var categories []domain.ProductCategory

	err := r.db.Model(&domain.ProductCategory{}).Find(&categories).Error
	if err != nil {
		return nil, err
	}

	return categories, nil
}
