package handler

import (
	"context"
	"errors"
	"time"

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
	var productIDs []string

	for _, v := range payload.Orders {
		if len(v.ProductID) == 0 {
			log.Error("invalid param: missing product id")
			return errors.New("invalid param: missing product id")
		}
		if v.Quantity < 1 {
			log.Error("invalid param: missing quantity or must be greater than 0 ")
			return errors.New("invalid param: missing quantity or must be greater than 0 ")
		}
		productIDs = append(productIDs, v.ProductID)
	}

	order := domain.Order{
		UserID:    userID,
		SubTotal:  0,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	err = h.repository.SaveOrder(ctx, order, payload.Orders, productIDs)
	if err != nil {
		log.Error("fail to create order: ", err.Error())
		return
	}

	return
}
