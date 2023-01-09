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

func (r *ProductRepository) GetProductList(ctx context.Context, page int, category string) ([]domain.ProductDTO, error) {
	var products []domain.ProductDTO

	modelQuery := r.db.Limit(50).Model(&domain.Product{}).Preload("ProductImage").Where("product_categories.category = ?", category)
	if len(category) == 0 {
		modelQuery = r.db.Limit(50).Model(&domain.Product{}).Preload("ProductImage")
	}

	modelQuery.Select("products.id, products.product_name, product_categories.category, products.price, products.available, products.description, products.min_order").Joins("left join product_categories on product_categories.id = products.category_id").Scan(&products)

	return products, nil
}
