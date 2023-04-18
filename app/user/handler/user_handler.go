package handler

import (
	"context"
	"time"

	"github.com/ryanadiputraa/api-udrio/domain"
	log "github.com/sirupsen/logrus"
)

type userHandler struct {
	repository     domain.IUserRepository
	cartRepository domain.ICartRepository
}

func NewUserHandler(repository domain.IUserRepository, cartRepository domain.ICartRepository) domain.IUserHandler {
	return &userHandler{repository: repository, cartRepository: cartRepository}
}

func (h *userHandler) CreateOrUpdateIfExist(ctx context.Context, user domain.User) error {
	user.CreatedAt = time.Now().UTC()
	user.UpdatedAt = time.Now().UTC()

	err := h.repository.SaveOrUpdate(ctx, user)
	if err != nil {
		log.Error("fail to save or update user: ", err.Error())
		return err
	}

	cart := domain.Cart{
		UserID:    user.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = h.cartRepository.CreateOrUpdate(ctx, cart)
	if err != nil {
		log.Error("fail create user cart: ", err.Error())
		return err
	}

	return nil
}

func (h *userHandler) GetUserInfo(ctx context.Context, userID string) (user domain.User, err error) {
	user, err = h.repository.FindByID(ctx, userID)
	return
}
