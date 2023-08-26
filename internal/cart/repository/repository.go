package repository

import (
	"context"

	"github.com/ryanadiputraa/api-udrio/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type repository struct {
	db *gorm.DB
}

func NewCartRepository(conn *gorm.DB) domain.CartRepository {
	return &repository{
		db: conn,
	}
}

func (r *repository) CreateOrUpdate(ctx context.Context, cart domain.Cart) error {
	err := r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"updated_at"}),
	}).Create(&cart).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) FetchCartByUserID(ctx context.Context, userID string) (cart []domain.CartDTO, err error) {
	err = r.db.Model(&domain.Cart{}).Select("cart_items.quantity, cart_items.created_at, products.id AS product_id, products.product_name, products.price, products.is_available, products.description, products.min_order, product_categories.category, product_categories.icon").Joins("LEFT JOIN cart_items ON carts.id = cart_items.cart_id LEFT JOIN products ON cart_items.product_id = products.id LEFT JOIN product_categories ON products.product_category_id = product_categories.id").Order("created_at desc").Where(&domain.Cart{UserID: userID}).Scan(&cart).Error

	if err != nil {
		return cart, err
	}

	for i, c := range cart {
		// handle 0 cart items
		if i == 0 && c.ProductID == "" {
			cart = []domain.CartDTO{}
			return
		}

		type ProductImg struct {
			Image string
		}
		var img ProductImg

		r.db.Model(&domain.ProductImage{}).Select("image").Where(&domain.ProductImage{ProductID: c.ProductID}).First(&img)
		cart[i].Image = img.Image
	}

	return cart, nil
}

func (r *repository) FindUserCartID(ctx context.Context, userID string) (cartID int, err error) {
	err = r.db.Model(&domain.Cart{}).Select("id").Where(&domain.Cart{UserID: userID}).Find(&cartID).Error
	if err != nil {
		return cartID, err
	}
	return cartID, nil
}

func (r *repository) PatchUserCart(ctx context.Context, cartItem domain.CartItem) error {
	err := r.db.Where(&domain.CartItem{CartID: cartItem.CartID}).Clauses(
		clause.OnConflict{
			Columns:   []clause.Column{{Name: "product_id"}},
			DoUpdates: clause.AssignmentColumns([]string{"quantity"}),
		}).Create(&cartItem).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) DeleteCartItemByProductID(ctx context.Context, cartID int, productID string) (cartItem domain.CartItem, err error) {
	err = r.db.Clauses(clause.Returning{Columns: []clause.Column{{Name: "cart_id"}}}).Where(&domain.CartItem{CartID: cartID, ProductID: productID}).Delete(&cartItem).Error
	return
}
