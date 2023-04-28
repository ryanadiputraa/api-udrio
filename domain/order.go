package domain

import (
	"context"
	"time"

	"github.com/ryanadiputraa/api-udrio/pkg/pagination"
)

type IOrderRepository interface {
	FetchOrdersByUserID(ctx context.Context, userID string, size int, offset int) (orders []OrderDTO, count int64, err error)
	SaveOrder(ctx context.Context, order Order, items []OrderPayloadItem, productIDs []string) (user User, err error)
}

type IOrderHandler interface {
	GetUserOrders(ctx context.Context, userID string, size int, page int) (orders []OrderDTO, meta pagination.Page, err error)
	CreateOrder(ctx context.Context, userID string, payload OrderPayload) error
}

type Order struct {
	ID        int64  `gorm:"primaryKey;type:bigserial"`
	UserID    string `gorm:"index;not null"`
	User      User
	Items     []OrderItem
	SubTotal  int       `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

type OrderItem struct {
	ID         string `gorm:"primaryKey"`
	OrderID    int64  `gorm:"index"`
	Order      Order
	ProductID  string
	Product    Product
	Quantity   int `gorm:"not null"`
	Price      int `gorm:"not null"`
	TotalPrice int `gorm:"not null"`
}
type OrderDTO struct {
	ID        int64          `json:"order_id"`
	SubTotal  int            `json:"sub_total"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	Items     []OrderItemDTO `json:"items"`
}

type OrderItemDTO struct {
	ProductID   string   `json:"product_id"`
	ProductName string   `json:"product_name"`
	Quantity    int      `json:"quantity"`
	Price       int      `json:"price"`
	TotalPrice  int      `json:"total_price"`
	Image       string   `json:"image"`
	OrderDTOID  string   `json:"-"`
	OrderDTO    OrderDTO `json:"-"`
}

type OrderPayload struct {
	Orders []OrderPayloadItem `json:"orders"`
}

type OrderPayloadItem struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}
