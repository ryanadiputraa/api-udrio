package handler

import (
	"context"

	"github.com/ryanadiputraa/api-udrio/domain"
)

type orderHandler struct {
	repository domain.IOrderRepository
}

func NewOrderHandler(repository domain.IOrderRepository) domain.IOrderHandler {
	return &orderHandler{repository: repository}
}

func (h *orderHandler) GetUserOrders(ctx context.Context, userID string) (orders []domain.OrderDTO, err error) {
	orders, err = h.repository.FetchOrdersByUserID(ctx, userID)
	if orders == nil {
		orders = []domain.OrderDTO{}
	}
	return
}

func (h *orderHandler) CreateOrder(ctx context.Context, userID string, payload domain.OrderPayload) (err error) {
	return
}
