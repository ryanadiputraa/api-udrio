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

func (r *ProductRepository) GetProductList(ctx context.Context, page int, category int) ([]domain.ProductDTO, error) {
	var products []domain.ProductDTO

	modelQuery := r.db.Limit(50).Model(&domain.Product{}).Where("products.category_id = ?", category)
	if category == 0 {
		modelQuery = r.db.Limit(50).Model(&domain.Product{})
	}

	modelQuery.Select("products.id, products.product_name, products.category_id, products.price, products.available, products.description, products.min_order").Scan(&products)

	return products, nil
}
