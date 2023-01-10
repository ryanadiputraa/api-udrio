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

	rows, err := modelQuery.Limit(50).Order("available desc, updated_at desc, created_at desc").Select("products.id, products.product_name, products.category_id, products.price, products.available, products.description, products.min_order").Rows()

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product domain.Product
		var productDTO domain.ProductDTO
		r.db.ScanRows(rows, &product)

		productImages, err := r.GetProductImages(ctx, product.ID)
		if err != nil {
			return nil, err
		}
		productDTO.ID = product.ID
		productDTO.ProductName = product.ProductName
		productDTO.CategoryID = product.CategoryID
		productDTO.Price = product.Price
		productDTO.Available = product.Available
		productDTO.Description = product.Description
		productDTO.MinOrder = product.MinOrder
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
