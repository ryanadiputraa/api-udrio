package repository

import (
	"context"
	"encoding/json"

	"github.com/ryanadiputraa/api-udrio/domain"
	"github.com/ryanadiputraa/api-udrio/pkg/cache"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type userRepository struct {
	db    *gorm.DB
	redis cache.Redis
}

func NewUserRepository(conn *gorm.DB, redis cache.Redis) domain.IUserRepository {
	return &userRepository{db: conn, redis: redis}
}

func (r *userRepository) SaveOrUpdate(ctx context.Context, user domain.User) error {
	err := r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"first_name", "last_name", "email", "picture", "locale", "updated_at"}),
	}).Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) FindByID(ctx context.Context, userID interface{}) (user domain.User, err error) {
	val, err := r.redis.Get(ctx, userID.(string))
	if err != nil {
		err = r.db.First(&user, "id = ?", userID).Error
		if err != nil {
			return
		}

		err = r.redis.Set(ctx, user.ID, user)
		return
	}

	err = json.Unmarshal([]byte(val), &user)
	return
}
