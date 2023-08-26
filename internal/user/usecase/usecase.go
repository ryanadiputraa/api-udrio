package usecase

import (
	"context"
	"time"

	"github.com/ryanadiputraa/api-udrio/domain"
)

type usecase struct {
	repository     domain.UserRepository
	cartRepository domain.CartRepository
}

func NewUserUsecase(repository domain.UserRepository, cartRepository domain.CartRepository) domain.UserUsecase {
	return &usecase{repository: repository, cartRepository: cartRepository}
}

func (u *usecase) CreateOrUpdateIfExist(ctx context.Context, user domain.User) error {
	user.CreatedAt = time.Now().UTC()
	user.UpdatedAt = time.Now().UTC()

	err := u.repository.SaveOrUpdate(ctx, user)
	if err != nil {
		return err
	}

	cart := domain.Cart{
		UserID:    user.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = u.cartRepository.CreateOrUpdate(ctx, cart)
	if err != nil {
		return err
	}

	return nil
}

func (u *usecase) GetUserInfo(ctx context.Context, userID string) (user domain.User, err error) {
	user, err = u.repository.FindByID(ctx, userID)
	return
}
