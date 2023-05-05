package handler

import (
	"context"
	"encoding/csv"
	"errors"
	"os"
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
		if err.Error() == "record not found" {
			err = errors.New("username tidak ditemukan")
		}
		log.Error("fail to signin: ", err.Error())
		return
	}

	if !utils.CheckPasswordWithHash(password, admin.Password) {
		log.Error("fail to signin: password didn't match")
		return "", time.Now(), errors.New("gagal login: password tidak sesuai")
	}

	sessionToken = uuid.NewString()
	expiresAt = time.Now().Add(time.Hour)
	expiresDuration := time.Hour
	session := domain.Session{
		SessionToken: sessionToken,
		ID:           admin.ID,
		Username:     admin.Username,
	}
	err = h.repository.SaveSession(ctx, session, expiresDuration)

	return
}

func (h *adminHandler) GetSession(ctx context.Context, sessionToken string) (session domain.Session, err error) {
	session, err = h.repository.GetSession(ctx, sessionToken)
	return
}

func (h *adminHandler) SaveFilePath(ctx context.Context, assetsPath domain.AssetsPath) (err error) {
	oldPath, err := h.repository.GetFilePath(ctx, assetsPath.Key)
	if err != nil && err.Error() != "record not found" {
		log.Error("fail to get file path: ", err.Error())
		return
	}
	if oldPath.FilePath != "" {
		if err = os.Remove(oldPath.FilePath); err != nil {
			log.Error("fail to remove old path: ", err.Error())
			return
		}
	}

	if err = h.repository.SaveFilePath(ctx, assetsPath); err != nil {
		log.Error("fail to save file path: ", err.Error())
	}
	return
}

func (h *adminHandler) GetFilePath(ctx context.Context, key string) (assetsPath domain.AssetsPath, err error) {
	assetsPath, err = h.repository.GetFilePath(ctx, key)
	return
}

func (h *adminHandler) BulkInsertProducts(ctx context.Context) (err error) {
	path, err := h.GetFilePath(ctx, "products")
	if err != nil {
		log.Error("fail to get file path: ", err.Error())
		return
	}

	f, err := os.Open(path.FilePath)
	if err != nil {
		log.Error("fail to open csv: ", err.Error())
		return
	}

	cr := csv.NewReader(f)
	if err = h.repository.BulkInsertProducts(ctx, cr); err != nil {
		log.Error("fail to insert data: ", err.Error())
		return
	}
	return
}
