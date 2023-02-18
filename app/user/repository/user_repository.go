package repository

import (
	"context"

	"github.com/ryanadiputraa/api-udrio/domain"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(conn *gorm.DB) domain.IUserRepository {
	return &userRepository{db: conn}
}

func (r *userRepository) SaveOrUpdate(ctx context.Context, user domain.User) error {
	return nil
}
