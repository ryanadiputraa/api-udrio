package repository

import (
	"context"

	"github.com/ryanadiputraa/api-udrio/internal/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type repository struct {
	db *gorm.DB
}

func NewUserRepository(conn *gorm.DB) domain.UserRepository {
	return &repository{db: conn}
}

func (r *repository) SaveOrUpdate(ctx context.Context, user domain.User) (err error) {
	err = r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"first_name", "last_name", "email", "picture", "locale", "updated_at"}),
	}).Create(&user).Error
	return
}

func (r *repository) FindByID(ctx context.Context, userID string) (user domain.User, err error) {
	err = r.db.First(&user, "id = ?", userID).Error
	return
}
