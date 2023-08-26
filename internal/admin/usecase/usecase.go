package usecase

import (
	"context"
	"encoding/csv"
	"errors"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/ryanadiputraa/api-udrio/internal/domain"
	"github.com/ryanadiputraa/api-udrio/pkg/utils"
)

type usecase struct {
	repository domain.AdminRepository
}

func NewAdminUsecase(repository domain.AdminRepository) domain.AdminUsecase {
	return &usecase{repository: repository}
}

func (u *usecase) SignIn(ctx context.Context, username string, password string) (sessionToken string, expiresAt time.Time, err error) {
	admin, err := u.repository.GetAdminByUsername(ctx, username)
	if err != nil {
		if err.Error() == "record not found" {
			err = errors.New("username tidak ditemukan")
		}
		return
	}

	if !utils.CheckPasswordWithHash(password, admin.Password) {
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
	err = u.repository.SaveSession(ctx, session, expiresDuration)

	return
}

func (u *usecase) GetSession(ctx context.Context, sessionToken string) (session domain.Session, err error) {
	session, err = u.repository.GetSession(ctx, sessionToken)
	return
}

func (u *usecase) SaveFilePath(ctx context.Context, assetsPath domain.AssetsPath) (err error) {
	oldPath, err := u.repository.GetFilePath(ctx, assetsPath.Key)
	if err != nil && err.Error() != "record not found" {
		return
	}
	if oldPath.FilePath != "" {
		if err = os.Remove(oldPath.FilePath); err != nil {
			return
		}
	}

	if err = u.repository.SaveFilePath(ctx, assetsPath); err != nil {
	}
	return
}

func (u *usecase) GetFilePath(ctx context.Context, key string) (assetsPath domain.AssetsPath, err error) {
	assetsPath, err = u.repository.GetFilePath(ctx, key)
	return
}

func (u *usecase) BulkInsertProducts(ctx context.Context) (err error) {
	path, err := u.GetFilePath(ctx, "products")
	if err != nil {
		return
	}

	f, err := os.Open(path.FilePath)
	if err != nil {
		return
	}

	cr := csv.NewReader(f)
	if err = u.repository.BulkInsertProducts(ctx, cr); err != nil {
		return
	}
	return
}
