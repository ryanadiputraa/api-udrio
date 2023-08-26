package domain

import (
	"context"
	"time"
)

type UserRepository interface {
	SaveOrUpdate(ctx context.Context, user User) error
	FindByID(ctx context.Context, userID string) (User, error)
}

type UserUsecase interface {
	CreateOrUpdateIfExist(ctx context.Context, user User) error
	GetUserInfo(ctx context.Context, userID string) (User, error)
}

type User struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	FirstName string    `gorm:"not null;type:varchar(50)" json:"first_name"`
	LastName  string    `gorm:"not null;type:varchar(50)" json:"last_name"`
	Email     string    `gorm:"unique;not null;type:varchar(100)" json:"email"`
	Picture   string    `gorm:"type:varchar(256)" json:"picture"`
	Locale    string    `gorm:"type:varchar(30)" json:"locale"`
	CreatedAt time.Time `gorm:"not null" json:"-"`
	UpdatedAt time.Time `gorm:"not null" json:"-"`
}
