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

	modelQuery := r.db.Model(&domain.Product{}).Where("products.category_id = ?", category)
	if category == 0 {
		modelQuery = r.db.Model(&domain.Product{})
	}

	rows, err := modelQuery.Limit(50).Order("available desc, updated_at desc, created_at desc").Select("products.id, products.product_name, products.category_id, product_categories.category, products.price, products.available, products.description, products.min_order").Joins("left join product_categories on product_categories.category_id = products.category_id").Rows()

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var productDTO domain.ProductDTO
		r.db.ScanRows(rows, &productDTO)

		productImages, err := r.GetProductImages(ctx, productDTO.ID)
		if err != nil {
			return nil, err
		}
		productDTO.Images = productImages
		products = append(products, productDTO)
	}

	return products, nil
}

func (r *ProductRepository) GetProductImages(ctx context.Context, productID string) ([]string, error) {
	var productImages []string

	rows, err := r.db.Model(&domain.ProductImage{}).Where("product_images.product_id = ?", productID).Select("product_images.image").Rows()

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var image domain.ProductImage
		r.db.ScanRows(rows, &image)
		productImages = append(productImages, image.Image)
	}

	if productImages == nil {
		return []string{}, nil
	}
	return productImages, nil
}
