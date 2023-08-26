package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"strings"
	"time"

	"cloud.google.com/go/storage"
	"github.com/ryanadiputraa/api-udrio/internal/domain"
	"github.com/ryanadiputraa/api-udrio/pkg/database"
	"gorm.io/gorm"
)

type repository struct {
	db      *gorm.DB
	redis   database.Redis
	storage *storage.BucketHandle
}

func NewProductRepository(conn *gorm.DB, redis database.Redis, storage *storage.BucketHandle) domain.ProductRepository {
	return &repository{db: conn, redis: redis, storage: storage}
}

func (r *repository) Fetch(ctx context.Context, size int, offset int, category int, query string) (products []domain.Product, count int64, err error) {
	modelQuery := r.db.Model(&domain.Product{}).Joins("ProductCategory", r.db.Where(&domain.ProductCategory{ID: category}))
	if category == 0 {
		modelQuery = r.db.Model(&domain.Product{}).Joins("ProductCategory")
		if len(query) != 0 {
			// using :* for half queries and & for multiple words
			query = strings.Replace(query, " ", ":*&", -1) + ":*"
			modelQuery = modelQuery.Where("to_tsvector(product_name) @@ to_tsquery(?)", query)
		}
	}

	modelQuery.Count(&count)

	err = modelQuery.Preload("ProductCategory").Preload("ProductImages").
		Limit(size).Offset(offset).Order("is_available desc, updated_at desc, created_at desc").
		Find(&products).Error
	if err != nil {
		return nil, 0, err
	}

	return products, count, nil
}

func (r *repository) FindByID(ctx context.Context, productID string) (domain.Product, error) {
	var product domain.Product

	err := r.db.Model(&domain.Product{}).Joins("ProductCategory").
		Where(&domain.Product{ID: productID}).Preload("ProductImages").
		First(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repository) FetchCategory(ctx context.Context) ([]domain.ProductCategory, error) {
	var categories []domain.ProductCategory
	cache, err := r.redis.Get(ctx, "products:category")
	if err != nil {
		err := r.db.Order("category").Find(&categories).Error
		if err != nil {
			return nil, err
		}

		err = r.redis.Set(ctx, "products:category", categories, time.Minute*10)
		if err != nil {
			return nil, err
		}
		return categories, nil
	}

	err = json.Unmarshal([]byte(cache), &categories)

	return categories, err
}

func (r *repository) SaveImage(ctx context.Context, file []byte, image domain.ProductImage) (err error) {
	err = r.db.Transaction(func(tx *gorm.DB) error {
		if err = r.db.Create(&image).Error; err != nil {
			return err
		}

		sw := r.storage.Object(image.ID).NewWriter(ctx)
		defer func() {
			if err = sw.Close(); err != nil {
				return
			}
		}()

		if _, err = io.Copy(sw, bytes.NewReader(file)); err != nil {
			return err
		}

		return nil
	})

	return
}
