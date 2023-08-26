package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/ryanadiputraa/api-udrio/internal/domain"
)

type usecase struct {
	repository domain.CartRepository
}

func NewCartUsecase(repository domain.CartRepository) domain.CartUsecase {
	return &usecase{
		repository: repository,
	}
}

func (u *usecase) GetUserCart(ctx context.Context, userID string) (cart []domain.CartDTO, err error) {
	cart, err = u.repository.FetchCartByUserID(ctx, userID)
	if err != nil {
		return cart, err
	}
	if cart == nil {
		cart = []domain.CartDTO{}
	}

	return cart, nil
}

func (u *usecase) UpdateUserCart(ctx context.Context, userID string, payload domain.CartPayload) error {
	cartID, err := u.repository.FindUserCartID(ctx, userID)
	if err != nil {
		return err
	}
	if cartID == 0 {
		return errors.New("cart not found")
	}

	cartItem := domain.CartItem{
		ID:        uuid.NewString(),
		CartID:    cartID,
		ProductID: payload.ProductID,
		Quantity:  payload.Quantity,
		CreatedAt: time.Now(),
	}
	err = u.repository.PatchUserCart(ctx, cartItem)
	if err != nil {
		return err
	}

	return nil
}

func (u *usecase) DeleteCartItem(ctx context.Context, userID string, productID string) (err error) {
	cartID, err := u.repository.FindUserCartID(ctx, userID)
	if err != nil {
		return
	}

	cartItem, err := u.repository.DeleteCartItemByProductID(ctx, cartID, productID)
	if err != nil {
		return
	}
	if cartItem.CartID == 0 {
		return errors.New("invalid param: product_id")
	}

	return
}
