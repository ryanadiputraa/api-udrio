package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/ryanadiputraa/api-udrio/domain"
	"github.com/ryanadiputraa/api-udrio/pkg/database"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type productRepository struct {
	db    *gorm.DB
	redis database.Redis
}

func NewProductRepository(conn *gorm.DB, redis database.Redis) domain.IProductRepository {
	return &productRepository{db: conn, redis: redis}
}

func (r *productRepository) Fetch(ctx context.Context, size int, offset int, category int, query string) (products []domain.Product, count int64, err error) {
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

	categoryStr := strconv.Itoa(category)
	sizeStr := strconv.Itoa(size)
	offsetStr := strconv.Itoa(offset)
	redisKey := fmt.Sprintf("products#%s#%s#%s#%s", categoryStr, sizeStr, offsetStr, query)

	cache, err := r.redis.Get(ctx, redisKey)
	if err != nil {
		err = modelQuery.Preload("ProductCategory").Preload("ProductImages").
			Limit(size).Offset(offset).Order("is_available desc, updated_at desc, created_at desc").
			Find(&products).Error
		if err != nil {
			return nil, 0, err
		}

		err = r.redis.Set(ctx, redisKey, products, time.Minute*10)
		if err != nil {
			return nil, 0, err
		}
		return products, count, nil
	}

	err = json.Unmarshal([]byte(cache), &products)
	if err != nil {
		return nil, 0, err
	}

	return products, count, nil
}

func (r *productRepository) FindByID(ctx context.Context, productID string) (domain.Product, error) {
	var product domain.Product

	err := r.db.Model(&domain.Product{}).Joins("ProductCategory").
		Where(&domain.Product{ID: productID}).Preload("ProductImages").
		First(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *productRepository) FetchCategory(ctx context.Context) ([]domain.ProductCategory, error) {
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

func (r *productRepository) UploadImage(ctx context.Context, file []byte, filename string) (url string, err error) {
	sw := database.FirebaseBucket.Object(filename).NewWriter(ctx)
	defer func() {
		err = sw.Close()
		if err != nil {
			return
		}
	}()

	if _, err = io.Copy(sw, bytes.NewReader(file)); err != nil {
		return
	}

	bucketName := viper.GetString("FIREBASE_BUCKET")
	storageToken := viper.GetString("FIREBASE_STORAGE_TOKEN")
	url = "https://firebasestorage.googleapis.com/v0/b/" + bucketName + "/o/" + filename + "?alt=media&token=" + storageToken

	return
}
