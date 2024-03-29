package repository

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"errors"
	"io"
	"strconv"
	"time"

	"github.com/ryanadiputraa/api-udrio/internal/domain"
	"github.com/ryanadiputraa/api-udrio/pkg/database"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type repository struct {
	db    *gorm.DB
	redis database.Redis
}

func NewAdminRepository(db *gorm.DB, redis database.Redis) domain.AdminRepository {
	return &repository{db: db, redis: redis}
}

func (r *repository) GetAdminByUsername(ctx context.Context, username string) (admin domain.Admin, err error) {
	err = r.db.First(&admin, "username = ?", username).Error
	return
}

func (r *repository) GetAdminByID(ctx context.Context, ID int) (admin domain.Admin, err error) {
	err = r.db.First(&admin, "id = ?", ID).Error
	return
}

func (r *repository) SaveSession(ctx context.Context, session domain.Session, expiresDuration time.Duration) (err error) {
	err = r.redis.Set(ctx, session.SessionToken, session, expiresDuration)
	return
}

func (r *repository) GetSession(ctx context.Context, sessionToken string) (session domain.Session, err error) {
	value, err := r.redis.Get(ctx, sessionToken)
	if err != nil {
		return
	}

	err = json.Unmarshal([]byte(value), &session)
	return
}

func (r *repository) SaveFilePath(ctx context.Context, assetsPath domain.AssetsPath) (err error) {
	err = r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "key"}},
		DoUpdates: clause.AssignmentColumns([]string{"file_path"}),
	}).Create(&assetsPath).Error
	return
}

func (r *repository) GetFilePath(ctx context.Context, key string) (assetsPath domain.AssetsPath, err error) {
	err = r.db.First(&assetsPath, "key = ?", key).Error
	return
}

func (r *repository) BulkInsertProducts(ctx context.Context, cr *csv.Reader) (err error) {
	var products []domain.Product
	header := true

	for {
		record, err := cr.Read()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return err
		}

		if !header {
			categoryID, _ := strconv.Atoi(record[2])
			price, _ := strconv.Atoi(record[3])
			isAvailable, _ := strconv.ParseBool(record[4])
			minOrder, _ := strconv.Atoi(record[6])
			product := domain.Product{
				ID:                record[0],
				ProductName:       record[1],
				ProductCategoryID: categoryID,
				Price:             price,
				IsAvailable:       isAvailable,
				Description:       record[5],
				MinOrder:          minOrder,
				CreatedAt:         time.Now(),
				UpdatedAt:         time.Now(),
			}

			products = append(products, product)
		}
		header = false
	}

	err = r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"product_name", "product_category_id", "price", "is_available", "description", "min_order", "updated_at"}),
	}).Create(&products).Error

	return
}
