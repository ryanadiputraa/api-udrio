package repository

import (
	"context"

	"github.com/ryanadiputraa/api-udrio/domain"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type cartRepository struct {
	db *gorm.DB
}

func NewCartRepository(conn *gorm.DB) domain.ICartRepository {
	return &cartRepository{
		db: conn,
	}
}

func (r *cartRepository) CreateOrUpdate(ctx context.Context, cart domain.Cart) error {
	err := r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"updated_at"}),
	}).Create(&cart).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *cartRepository) FetchCartByUserID(ctx context.Context, userID string) (cart []domain.CartDTO, err error) {
	err = r.db.Model(&domain.Cart{}).Select("cart_items.quantity, products.id AS product_id, products.product_name, products.price, products.is_available, products.description, products.min_order, product_categories.category, product_categories.icon").Joins("LEFT JOIN cart_items ON carts.id = cart_items.cart_id LEFT JOIN products ON cart_items.product_id = products.id LEFT JOIN product_categories ON products.product_category_id = product_categories.id").Where(&domain.Cart{UserID: userID}).Scan(&cart).Error

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

		err = r.db.Model(&domain.ProductImage{}).Select("image").Where(&domain.ProductImage{ProductID: c.ProductID}).First(&img).Error
		if err != nil {
			return cart, err
		}

		cart[i].Image = img.Image
	}

	return cart, nil
}

func (r *cartRepository) FindUserCartID(ctx context.Context, userID string) (cartID int, err error) {
	err = r.db.Model(&domain.Cart{}).Select("id").Where(&domain.Cart{UserID: userID}).Find(&cartID).Error
	if err != nil {
		return cartID, err
	}
	logrus.Error(cartID)
	return cartID, nil
}

func (r *cartRepository) PatchUserCart(ctx context.Context, cartItem domain.CartItem) error {
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

func (r *cartRepository) DeleteCartItemByProductID(ctx context.Context, cartID int, productID string) (cartItem domain.CartItem, err error) {
	err = r.db.Clauses(clause.Returning{Columns: []clause.Column{{Name: "cart_id"}}}).Where(&domain.CartItem{CartID: cartID, ProductID: productID}).Delete(&cartItem).Error
	return
}
