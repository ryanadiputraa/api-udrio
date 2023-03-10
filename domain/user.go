package domain

import (
	"context"
	"time"
)

type IUserRepository interface {
	SaveOrUpdate(ctx context.Context, user User) error
}

type IUserHandler interface {
	CreateOrUpdateIfExist(ctx context.Context, user User) error
	GetUserInfo(ctx context.Context, userID interface{}) (User, error)
}

type User struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"not null;type:varchar(100)" json:"name"`
	Email     string    `gorm:"unique;not null;type:varchar(100)" json:"email"`
	Picture   string    `gorm:"type:varchar(256)" json:"picture"`
	Locale    string    `gorm:"type:varchar(30)" json:"locale"`
	CreatedAt time.Time `gorm:"not null" json:"-"`
	UpdatedAt time.Time `gorm:"not null" json:"-"`
}
