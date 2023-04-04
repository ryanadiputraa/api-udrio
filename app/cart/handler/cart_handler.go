package handler

import (
	"context"
	"errors"

	"github.com/google/uuid"
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

func (h *cartHandler) UpdateUserCart(ctx context.Context, userID string, payload domain.CartPayload) error {
	cartID, err := h.repository.FindUserCartID(ctx, userID)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	if cartID == 0 {
		log.Error("cart not found")
		return errors.New("cart not found")
	}

	cartItem := domain.CartItem{
		ID:        uuid.NewString(),
		CartID:    cartID,
		ProductID: payload.ProductID,
		Quantity:  payload.Quantity,
	}
	err = h.repository.PatchUserCart(ctx, cartItem)
	if err != nil {
		log.Error("fail to update cart: ", err.Error())
		return err
	}

	return nil
}

func (h *cartHandler) DeleteCartItem(ctx context.Context, userID string, productID string) (err error) {
	cartID, err := h.repository.FindUserCartID(ctx, userID)
	if err != nil {
		log.Error(err.Error())
		return
	}

	cartItem, err := h.repository.DeleteCartItemByProductID(ctx, cartID, productID)
	if err != nil {
		log.Error(err.Error())
		return
	}
	if cartItem.CartID == 0 {
		log.Error("invalid param: product_id")
		return errors.New("invalid param: product_id")
	}

	return
}
