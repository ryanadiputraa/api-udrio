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
	ID              string          `gorm:"primaryKey"`
	ProductName     string          `gorm:"unique;not null;type:varchar(256)"`
	CategoryID      int             `gorm:"not null"`
	ProductCategory ProductCategory `gorm:"foreignKey:CategoryID;references:CategoryID"`
	Price           int             `gorm:"not null"`
	Available       bool            `gorm:"not null"`
	Description     string
	MinOrder        int            `gorm:"not null"`
	Images          []ProductImage `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt       time.Time      `gorm:"not null"`
	UpdatedAt       time.Time
}

type ProductDTO struct {
	ID          string    `json:"id"`
	ProductName string    `json:"product_name"`
	CategoryID  int       `json:"category_id"`
	Category    string    `json:"category"`
	Price       int       `json:"price"`
	Available   bool      `json:"available"`
	Description string    `json:"description"`
	MinOrder    int       `json:"min_order"`
	Images      []string  `json:"images" gorm:"type:text"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ProductImage struct {
	ID        string `gorm:"primaryKey"`
	Image     string `gorm:"not null;type:varchar(256)"`
	ProductID string `gorm:"foreignKey"`
}

type ProductCategory struct {
	CategoryID int    `gorm:"primaryKey;type:serial"`
	Category   string `gorm:"not null;unique;type:varchar(100)"`
	Icon       string `gorm:"not null;type:varchar(256)"`
}
