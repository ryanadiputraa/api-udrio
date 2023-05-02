package handler

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/ryanadiputraa/api-udrio/domain"
	"github.com/ryanadiputraa/api-udrio/pkg/utils"
	log "github.com/sirupsen/logrus"
)

type adminHandler struct {
	repository domain.IAdminRepository
}

func NewAdminHandler(repository domain.IAdminRepository) domain.IAdminHandler {
	return &adminHandler{repository: repository}
}

func (h *adminHandler) SignIn(ctx context.Context, username string, password string) (sessionToken string, expiresAt time.Time, err error) {
	admin, err := h.repository.GetAdminByUsername(ctx, username)
	if err != nil {
		log.Error("fail to signin: ", err.Error())
		return
	}

	if !utils.CheckPasswordWithHash(password, admin.Password) {
		log.Error("fail to signin: password didn't match")
		return "", time.Now(), errors.New("fail to signin: password didn't match")
	}

	sessionToken = uuid.NewString()
	expiresAt = time.Now().Add(time.Minute * 10)
	session := domain.Session{
		SessionToken: sessionToken,
		ID:           admin.ID,
		Username:     admin.Username,
	}
	err = h.repository.SaveSession(ctx, session)

	return
}

func (h *adminHandler) GetSession(ctx context.Context, sessionToken string) (session domain.Session, err error) {
	session, err = h.repository.GetSession(ctx, sessionToken)
	return
}
