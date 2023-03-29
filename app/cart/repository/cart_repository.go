package repository

import (
	"context"

	"github.com/ryanadiputraa/api-udrio/domain"
	"gorm.io/gorm"
)

type cartRepository struct {
	DB *gorm.DB
}

func NewCartRepository(conn *gorm.DB) domain.ICartRepository {
	return &cartRepository{
		DB: conn,
	}
}

func (r *cartRepository) FetchCartByUserID(ctx context.Context, userID string) (cart []domain.CartDTO, err error) {
	err = r.DB.Model(&domain.Cart{}).Select("cart_items.quantity, products.id AS product_id, products.product_name, products.price, products.is_available, products.description, products.min_order, product_categories.category, product_categories.icon").Joins("LEFT JOIN cart_items ON carts.id = cart_items.cart_id LEFT JOIN products ON cart_items.product_id = products.id LEFT JOIN product_categories ON products.product_category_id = product_categories.id").Where(&domain.Cart{UserID: userID}).Scan(&cart).Error

	if err != nil {
		return cart, nil
	}

	for i, c := range cart {
		type ProductImg struct {
			Image string
		}
		var img ProductImg

		err = r.DB.Model(&domain.ProductImage{}).Select("image").Where(&domain.ProductImage{ProductID: c.ProductID}).First(&img).Error
		if err != nil {
			return cart, nil
		}

		cart[i].Image = img.Image
	}

	return cart, nil
}
