package domain

import (
	"context"
	"encoding/csv"
	"time"
)

type AdminUsecase interface {
	SignIn(ctx context.Context, username string, password string) (sessionToken string, expiresAt time.Time, err error)
	GetSession(ctx context.Context, sessionToken string) (Session, error)
	SaveFilePath(ctx context.Context, assetsPath AssetsPath) error
	GetFilePath(ctx context.Context, key string) (AssetsPath, error)
	BulkInsertProducts(ctx context.Context) error
}

type AdminRepository interface {
	GetAdminByUsername(ctx context.Context, username string) (Admin, error)
	GetAdminByID(ctx context.Context, id int) (Admin, error)
	SaveSession(ctx context.Context, session Session, expiresDuration time.Duration) error
	GetSession(ctx context.Context, sessionToken string) (Session, error)
	SaveFilePath(ctx context.Context, assetsPath AssetsPath) error
	GetFilePath(ctx context.Context, key string) (AssetsPath, error)
	BulkInsertProducts(ctx context.Context, csvReader *csv.Reader) error
}

type Admin struct {
	ID       int    `gorm:"primaryKey;type:serial" json:"id"`
	Username string `gorm:"not null;unique;index" json:"username"`
	Password string `gorm:"not null" json:"passowrd"`
}

type Session struct {
	SessionToken string `json:"session_token"`
	ID           int    `json:"id"`
	Username     string `json:"username"`
}

type AssetsPath struct {
	Key      string `gorm:"primaryKey"`
	FilePath string `gorm:"not null"`
}
