package domain

import (
	"context"
	"time"

	"github.com/ryanadiputraa/api-udrio/pkg/pagination"
)

type IProductRepository interface {
	Fetch(ctx context.Context, limit int, offset int, categoryID int, query string) (products []Product, count int64, err error)
	FindByID(ctx context.Context, productID string) (Product, error)
	FetchCategory(ctx context.Context) ([]ProductCategory, error)
}

type IProductHandler interface {
	GetProductList(ctx context.Context, size int, page int, categoryID int, query string) (products []Product, meta pagination.Page, err error)
	GetProductDetail(ctx context.Context, productID string) (Product, error)
	GetProductCategoryList(ctx context.Context) ([]ProductCategory, error)
}

type Product struct {
	ID                string          `gorm:"primaryKey" json:"id"`
	ProductName       string          `gorm:"unique;index;not null;type:varchar(256)" json:"product_name"`
	ProductCategoryID int             `gorm:"index" json:"-"`
	ProductCategory   ProductCategory `json:"product_category"`
	Price             int             `gorm:"not null" json:"price"`
	IsAvailable       bool            `gorm:"not null" json:"is_available"`
	Description       string          `json:"description"`
	ProcessingTime    string          `json:"processing_time"`
	MinOrder          int             `gorm:"not null" json:"min_order"`
	ProductImages     []ProductImage  `json:"images"`
	CreatedAt         time.Time       `gorm:"not null" json:"created_at"`
	UpdatedAt         time.Time       `gorm:"not null" json:"updated_at"`
}

type ProductCategory struct {
	ID       int       `gorm:"primaryKey" json:"category_id"`
	Category string    `gorm:"not null;unique;type:varchar(100)" json:"category"`
	Icon     string    `gorm:"not null;type:varchar(256)" json:"icon"`
	Products []Product `json:"-"`
}

type ProductImage struct {
	ID        string  `gorm:"primaryKey" json:"image_id"`
	Image     string  `gorm:"not null;type:varchar(256)" json:"url"`
	ProductID string  `gorm:"index" json:"-"`
	Product   Product `json:"-"`
}
