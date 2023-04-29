package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/ryanadiputraa/api-udrio/domain"
	"github.com/ryanadiputraa/api-udrio/pkg/cache"
	"gorm.io/gorm"
)

type productRepository struct {
	db    *gorm.DB
	redis cache.Redis
}

func NewProductRepository(conn *gorm.DB, redis cache.Redis) domain.IProductRepository {
	return &productRepository{db: conn, redis: redis}
}

func (r *productRepository) Fetch(ctx context.Context, size int, offset int, category int, query string) (products []domain.Product, count int64, err error) {
	r.db.Model(&domain.Product{}).Count(&count)

	modelQuery := r.db.Model(&domain.Product{}).Joins("ProductCategory", r.db.Where(&domain.ProductCategory{ID: category}))
	if category == 0 {
		modelQuery = r.db.Model(&domain.Product{}).Joins("ProductCategory")
		if len(query) != 0 {
			// using :* for half queries and & for multiple words
			query = strings.Replace(query, " ", ":*&", -1) + ":*"
			modelQuery = modelQuery.Where("to_tsvector(product_name) @@ to_tsquery(?)", query)
		}
	}

	categoryStr := strconv.Itoa(category)
	sizeStr := strconv.Itoa(size)
	offsetStr := strconv.Itoa(offset)
	redisKey := fmt.Sprintf("products#%s#%s#%s#%s", categoryStr, sizeStr, offsetStr, query)

	cache, err := r.redis.Get(ctx, redisKey)
	if err != nil {
		err = modelQuery.Preload("ProductImages").Limit(size).Offset(offset).Order("is_available desc, updated_at desc, created_at desc").Find(&products).Error
		if err != nil {
			return nil, 0, err
		}

		err = r.redis.Set(ctx, redisKey, products)
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

	err := r.db.Model(&domain.Product{}).Joins("ProductCategory").Where(&domain.Product{ID: productID}).Preload("ProductImages").First(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *productRepository) FetchCategory(ctx context.Context) ([]domain.ProductCategory, error) {
	var categories []domain.ProductCategory
	cache, err := r.redis.Get(ctx, "products:category")
	if err != nil {
		err := r.db.Find(&categories).Error
		if err != nil {
			return nil, err
		}

		err = r.redis.Set(ctx, "products:category", categories)
		if err != nil {
			return nil, err
		}
		return categories, nil
	}

	err = json.Unmarshal([]byte(cache), &categories)

	return categories, err
}
