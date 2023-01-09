package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type IProductRepository interface {
	GetProductList(ctx context.Context, db *gorm.DB, page int) ([]Product, error)
}

type IProductService interface {
	GetProductList(ctx context.Context, page int) ([]Product, error)
}

type Product struct {
	ID          string `gorm:"primaryKey"`
	ProductName string `gorm:"unique;not null;type:varchar(256)"`
	Category    int
	Price       int            `gorm:"not null"`
	Available   bool           `gorm:"not null"`
	Description string         `gorm:"foreignKey;references:ProductCategory"`
	MinOrder    int            `gorm:"not null"`
	Images      []ProductImage `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt   time.Time      `gorm:"not null"`
	UpdatedAt   time.Time
}

type ProductImage struct {
	ID        string `gorm:"primaryKey"`
	Image     string `gorm:"not null;type:varchar(256)"`
	ProductID string
}

type ProductCategory struct {
	ID       int    `gorm:"primaryKey;type:serial"`
	Category string `gorm:"not null;type:varchar(100)"`
	Icon     string `gorm:"not null;type:varchar(256)"`
}
