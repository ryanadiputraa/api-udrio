package repository

import (
	"context"
	"encoding/json"
	"time"

	"github.com/ryanadiputraa/api-udrio/domain"
	"github.com/ryanadiputraa/api-udrio/pkg/cache"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type adminRepository struct {
	db    *gorm.DB
	redis cache.Redis
}

func NewAdminRepository(db *gorm.DB, redis cache.Redis) domain.IAdminRepository {
	return &adminRepository{db: db, redis: redis}
}

func (r *adminRepository) GetAdminByUsername(ctx context.Context, username string) (admin domain.Admin, err error) {
	err = r.db.First(&admin, "username = ?", username).Error
	return
}

func (r *adminRepository) GetAdminByID(ctx context.Context, ID int) (admin domain.Admin, err error) {
	err = r.db.First(&admin, "id = ?", ID).Error
	return
}

func (r *adminRepository) SaveSession(ctx context.Context, session domain.Session) (err error) {
	err = r.redis.Set(ctx, session.SessionToken, session, time.Hour)
	return
}

func (r *adminRepository) GetSession(ctx context.Context, sessionToken string) (session domain.Session, err error) {
	value, err := r.redis.Get(ctx, sessionToken)
	if err != nil {
		return
	}

	err = json.Unmarshal([]byte(value), &session)
	return
}

func (r *adminRepository) SaveFilePath(ctx context.Context, assetsPath domain.AssetsPath) (err error) {
	err = r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "key"}},
		DoUpdates: clause.AssignmentColumns([]string{"file_path"}),
	}).Create(&assetsPath).Error
	return
}

func (r *adminRepository) GetFilePath(ctx context.Context, key string) (assetsPath domain.AssetsPath, err error) {
	err = r.db.First(&assetsPath, "key = ?", key).Error
	return
}
