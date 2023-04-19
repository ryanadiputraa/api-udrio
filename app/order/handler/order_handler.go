package handler

import (
	"context"

	"github.com/ryanadiputraa/api-udrio/domain"
	log "github.com/sirupsen/logrus"
)

type orderHandler struct {
	repository domain.IOrderRepository
}

func NewOrderHandler(repository domain.IOrderRepository) domain.IOrderHandler {
	return &orderHandler{repository: repository}
}

func (h *orderHandler) GetUserOrders(ctx context.Context, userID string) (order []domain.OrderDTO, err error) {
	order, err = h.repository.FetchOrdersByUserID(ctx, userID)
	if order == nil {
		order = []domain.OrderDTO{}
	}
	if err != nil {
		if err.Error() == "record not found" {
			return []domain.OrderDTO{}, nil
		}
		log.Error("fail to fetch user orders: ", err.Error())
	}

	return
}

func (h *orderHandler) CreateOrder(ctx context.Context, userID string, payload domain.OrderPayload) (err error) {
	return
}
