package domain

import (
	"context"
	"time"
)

type CartRepository interface {
	CreateOrUpdate(ctx context.Context, cart Cart) error
	FetchCartByUserID(ctx context.Context, userID string) ([]CartDTO, error)
	FindUserCartID(ctx context.Context, userID string) (int, error)
	PatchUserCart(ctx context.Context, cartItem CartItem) error
	DeleteCartItemByProductID(ctx context.Context, cartID int, productID string) (cartItem CartItem, err error)
}

type CartUsecase interface {
	GetUserCart(ctx context.Context, userID string) ([]CartDTO, error)
	UpdateUserCart(ctx context.Context, userID string, payload CartPayload) error
	DeleteCartItem(ctx context.Context, userID string, productID string) error
}

type Cart struct {
	ID        int    `gorm:"primaryKey;serial"`
	UserID    string `gorm:"index;not null;unique"`
	User      User
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
	CartItems []CartItem
}

type CartItem struct {
	ID        string `gorm:"primaryKey"`
	Quantity  int    `gorm:"not null"`
	CartID    int    `gorm:"index"`
	Cart      Cart
	ProductID string `gorm:"index;not null;unique"`
	Product   Product
	CreatedAt time.Time `gorm:"not null"`
}

type CartDTO struct {
	Quantity    int       `json:"quantity"`
	ProductID   string    `json:"product_id"`
	ProductName string    `json:"product_name"`
	Price       int       `json:"price"`
	IsAvailable bool      `json:"is_available"`
	Image       string    `json:"image"`
	MinOrder    int       `json:"min_order"`
	CreatedAt   time.Time `json:"created_at"`
}

type CartPayload struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}
