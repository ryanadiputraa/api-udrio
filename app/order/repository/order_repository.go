package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/ryanadiputraa/api-udrio/domain"
	"gorm.io/gorm"
)

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(conn *gorm.DB) domain.IOrderRepository {
	return &orderRepository{db: conn}
}

func (r *orderRepository) FetchOrdersByUserID(ctx context.Context, userID string, size int, offset int) (orders []domain.OrderDTO, count int64, err error) {
	r.db.Model(&domain.Order{}).Where(&domain.Order{UserID: userID}).Count(&count)

	rows, err := r.db.Model(&[]domain.Order{}).Limit(size).Offset(offset).Where(&domain.Order{UserID: userID}).Order("created_at DESC").Rows()
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var order domain.OrderDTO
		r.db.ScanRows(rows, &order)

		// fetch order items
		itemRows, itemErr := r.db.Model(&[]domain.OrderItem{}).Where(&domain.OrderItem{OrderID: order.ID}).Rows()
		if itemErr != nil {
			return
		}
		defer itemRows.Close()

		var items []domain.OrderItemDTO
		for itemRows.Next() {
			var item domain.OrderItemDTO
			r.db.ScanRows(itemRows, &item)

			var product domain.Product
			err = r.db.Model(&domain.Product{}).Where(&domain.Product{ID: item.ProductID}).Preload("ProductImages").First(&product).Error
			if err != nil {
				return
			}

			item.ProductName = product.ProductName
			if len(product.ProductImages) > 0 {
				item.Image = product.ProductImages[0].Image
			}

			items = append(items, item)
		}

		if items == nil {
			order.Items = []domain.OrderItemDTO{}
		} else {
			order.Items = items
		}
		orders = append(orders, order)
	}

	return
}

func (r *orderRepository) SaveOrder(ctx context.Context, order domain.Order, items []domain.OrderPayloadItem, productIDs []string) (err error) {
	rows, err := r.db.Model(&domain.Product{}).Where("id IN ?", productIDs).Rows()
	if err != nil {
		return
	}
	defer rows.Close()

	idx := 0
	for rows.Next() {
		var product domain.Product
		r.db.ScanRows(rows, &product)
		totalPrice := product.Price * items[idx].Quantity

		orderItem := domain.OrderItem{
			ID:         uuid.NewString(),
			ProductID:  product.ID,
			Quantity:   items[idx].Quantity,
			Price:      product.Price,
			TotalPrice: totalPrice,
		}
		order.SubTotal += totalPrice
		order.Items = append(order.Items, orderItem)

		idx++
	}

	if err = r.db.Create(&order).Error; err != nil {
		return err
	}

	return
}
