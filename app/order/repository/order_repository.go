package repository

import (
	"context"

	"github.com/ryanadiputraa/api-udrio/domain"
	"gorm.io/gorm"
)

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(conn *gorm.DB) domain.IOrderRepository {
	return &orderRepository{db: conn}
}

func (r *orderRepository) FetchOrdersByUserID(ctx context.Context, userID string) (orders []domain.OrderDTO, err error) {
	var order domain.Order
	if err = r.db.First(&order, "user_id = ?", userID).Error; err != nil {
		return
	}

	rows, err := r.db.Model(&[]domain.OrderItem{}).Where(&domain.OrderItem{OrderID: order.ID}).Rows()
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var orderItem domain.OrderDTO
		r.db.ScanRows(rows, &orderItem)

		var product domain.Product
		err = r.db.Model(&domain.Product{}).Where(&domain.Product{ID: orderItem.ProductID}).Preload("ProductImages").First(&product).Error
		if err != nil {
			return
		}

		orderItem.ProductName = product.ProductName
		if len(product.ProductImages) > 0 {
			orderItem.Image = product.ProductImages[0].Image
		}

		orders = append(orders, orderItem)
	}
	return
}

func (r *orderRepository) SaveOrder(ctx context.Context, userID string, order domain.Order) (err error) {
	err = r.db.Create(&order).Error
	return
}
