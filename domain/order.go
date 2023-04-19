package domain

import (
	"context"
	"time"
)

type IOrderRepository interface {
	FetchOrdersByUserID(ctx context.Context, userID string) ([]OrderDTO, error)
	SaveOrder(ctx context.Context, userID string, order Order) error
}

type IOrderHandler interface {
	GetUserOrders(ctx context.Context, userID string) ([]OrderDTO, error)
	CreateOrder(ctx context.Context, userID string, payload OrderPayload) error
}

type Order struct {
	ID        string `gorm:"primaryKey"`
	UserID    string `gorm:"index"`
	User      User
	Items     []OrderItem
	SubTotal  int       `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

type OrderItem struct {
	ID         int64  `gorm:"primaryKey;type:bigserial"`
	OrderID    string `gorm:"index"`
	Order      Order
	ProductID  string
	Product    Product
	Quantity   int `gorm:"not null"`
	TotalPrice int `gorm:"not null"`
}
type OrderDTO struct {
	ID       string         `json:"order_id"`
	SubTotal int            `json:"sub_total"`
	Items    []OrderItemDTO `json:"items"`
}

type OrderItemDTO struct {
	ProductID   string   `json:"product_id"`
	ProductName string   `json:"product_name"`
	Quantity    int      `json:"quantity"`
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
