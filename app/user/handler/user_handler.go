package handler

import (
	"context"

	"github.com/ryanadiputraa/api-udrio/domain"
)

type userHandler struct {
	repository domain.IUserRepository
}

func NewUserHandler(repository domain.IUserRepository) domain.IUserHandler {
	return &userHandler{repository: repository}
}

func (h *userHandler) CreateOrUpdateIfExist(ctx context.Context, user domain.User) error {
	return nil
}
