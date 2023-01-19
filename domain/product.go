package domain

import (
	"context"
	"time"
)

type IProductRepository interface {
	GetProductList(ctx context.Context, page int, categoryID int) ([]Product, error)
}

type IProductService interface {
	GetProductList(ctx context.Context, page int, categoryID int) ([]Product, error)
}

type Product struct {
	ID                string          `gorm:"primaryKey" json:"id"`
	ProductName       string          `gorm:"unique;not null;type:varchar(256)" json:"product_name"`
	ProductCategoryID int             `gorm:"index" json:"-"`
	ProductCategory   ProductCategory `json:"product_category"`
	Price             int             `gorm:"not null" json:"price"`
	IsAvailable       bool            `gorm:"not null" json:"is_available"`
	Description       string          `json:"description"`
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
	ProductID string  `json:"-"`
	Product   Product `json:"-"`
}
