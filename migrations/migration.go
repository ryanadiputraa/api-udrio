package migrations

import (
	"github.com/ryanadiputraa/api-udrio/domain"
	"github.com/ryanadiputraa/api-udrio/pkg/database"
)

func Migrate() {
	db := database.GetConnection()
	db.AutoMigrate(&domain.Product{}, &domain.ProductImage{}, &domain.ProductCategory{})
}
