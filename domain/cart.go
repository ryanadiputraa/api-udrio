package domain

import (
	"context"
	"time"
)

type ICartRepository interface {
	CreateOrUpdate(ctx context.Context, cart Cart) error
	FetchCartByUserID(ctx context.Context, userID string) ([]CartDTO, error)
	FindUserCartID(ctx context.Context, userID string) (int, error)
	PatchUserCart(ctx context.Context, cartItem CartItem) error
}

type ICartHandler interface {
	GetUserCart(ctx context.Context, userID string) ([]CartDTO, error)
	UpdateUserCart(ctx context.Context, userID string, payload CartPayload) error
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
}

type CartDTO struct {
	Quantity    int    `json:"quantity"`
	ProductID   string `json:"product_id"`
	ProductName string `json:"product_name"`
	Price       int    `json:"price"`
	IsAvailable bool   `json:"is_available"`
	Image       string `json:"image"`
	MinOrder    int    `json:"min_order"`
}

type CartPayload struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}
