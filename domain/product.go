package domain

import (
	"context"
	"time"
)

type IProductRepository interface {
	GetProductList(ctx context.Context, page int, categoryID int) ([]ProductDTO, error)
	GetProductImages(ctx context.Context, productID string) ([]string, error)
}

type IProductService interface {
	GetProductList(ctx context.Context, page int, categoryID int) ([]ProductDTO, error)
}

type Product struct {
	ID          string         `gorm:"primaryKey" json:"id"`
	ProductName string         `gorm:"unique;not null;type:varchar(256)" json:"product_name"`
	CategoryID  int            `gorm:"foreignKey;references:ProductCategory" json:"category_id"`
	Price       int            `gorm:"not null" json:"price"`
	Available   bool           `gorm:"not null" json:"available"`
	Description string         `json:"description"`
	MinOrder    int            `gorm:"not null" json:"min_order"`
	Images      []ProductImage `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"images"`
	CreatedAt   time.Time      `gorm:"not null" json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

type ProductDTO struct {
	ID          string   `json:"id"`
	ProductName string   `json:"product_name"`
	CategoryID  int      `json:"category_id"`
	Price       int      `json:"price"`
	Available   bool     `json:"available"`
	Description string   `json:"description"`
	MinOrder    int      `json:"min_order"`
	Images      []string `json:"images"`
}

type ProductImage struct {
	ID        string `gorm:"primaryKey"`
	Image     string `gorm:"not null;type:varchar(256)"`
	ProductID string `gorm:"foreignKey"`
}

type ProductCategory struct {
	ID       int    `gorm:"primaryKey;type:serial"`
	Category string `gorm:"not null;type:varchar(100)"`
	Icon     string `gorm:"not null;type:varchar(256)"`
}
