package repository

import (
	"context"

	"github.com/ryanadiputraa/api-udrio/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(conn *gorm.DB) domain.IUserRepository {
	return &userRepository{db: conn}
}

func (r *userRepository) SaveOrUpdate(ctx context.Context, user domain.User) (err error) {
	err = r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"first_name", "last_name", "email", "picture", "locale", "updated_at"}),
	}).Create(&user).Error
	return
}

func (r *userRepository) FindByID(ctx context.Context, userID string) (user domain.User, err error) {
	err = r.db.First(&user, "id = ?", userID).Error
	return
}
