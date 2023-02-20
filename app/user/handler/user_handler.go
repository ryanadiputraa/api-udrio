package handler

import (
	"context"
	"time"

	"github.com/ryanadiputraa/api-udrio/domain"
	log "github.com/sirupsen/logrus"
)

type userHandler struct {
	repository domain.IUserRepository
}

func NewUserHandler(repository domain.IUserRepository) domain.IUserHandler {
	return &userHandler{repository: repository}
}

func (h *userHandler) CreateOrUpdateIfExist(ctx context.Context, user domain.User) error {
	user.CreatedAt = time.Now().UTC()
	user.UpdatedAt = time.Now().UTC()

	err := h.repository.SaveOrUpdate(ctx, user)
	if err != nil {
		log.Error("failed to save or update user: ", err.Error())
		return err
	}

	return nil
}

func (h *userHandler) GetUserInfo(ctx context.Context, userID interface{}) (user domain.User, err error) {
	user, err = h.repository.FindByID(ctx, userID)
	return
}
