package handler

import (
	"context"

	"github.com/ryanadiputraa/api-udrio/domain"
	log "github.com/sirupsen/logrus"
)

type cartHandler struct {
	repository domain.ICartRepository
}

func NewCartHandler(repository domain.ICartRepository) domain.ICartHandler {
	return &cartHandler{
		repository: repository,
	}
}

func (h *cartHandler) GetUserCart(ctx context.Context, userID string) (cart []domain.CartDTO, err error) {
	cart, err = h.repository.FetchCartByUserID(ctx, userID)
	if err != nil {
		log.Error("fail to fetch user cart:", err.Error())
		return cart, err
	}
	if cart == nil {
		cart = []domain.CartDTO{}
	}

	return cart, nil
}
