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

func (r *userRepository) SaveOrUpdate(ctx context.Context, user domain.User) error {
	err := r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"name", "email", "picture", "locale", "updated_at"}),
	}).Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) FindByID(ctx context.Context, userID interface{}) (user domain.User, err error) {
	err = r.db.First(&user, "id = ?", userID).Error
	return
}
