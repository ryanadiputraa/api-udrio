package database

import (
	"time"

	"github.com/ryanadiputraa/api-udrio/domain"
	"gorm.io/gorm"
)

type Sedeer struct {
	Sedeer interface{}
}

func registerSedeers(db *gorm.DB) []Sedeer {
	sedeers := append(getCategorySedeers(), getProductSedeers()...)
	sedeers = append(sedeers, GetOrdersSeeder()...)
	return sedeers
}

func DBSeed(db *gorm.DB) error {
	for _, sedeer := range registerSedeers(db) {
		err := db.Debug().FirstOrCreate(sedeer.Sedeer).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func getCategorySedeers() []Sedeer {
	return []Sedeer{
		{Sedeer: &domain.ProductCategory{
			ID:       1,
			Category: "Spanduk & Banner",
			Icon:     "https://www.riodigitalprint.com/liteprint/application/liteprint/rio/assets/product_category/dcc63fc4f4e0293feca821c7b3d4de6f.png"}},
		{Sedeer: &domain.ProductCategory{
			ID:       2,
			Category: "Cetak Digital & Buku",
			Icon:     "https://www.riodigitalprint.com/liteprint/application/liteprint/rio/assets/product_category/f40863821be94d257f959ccbb1b77f23.png"}},
		{Sedeer: &domain.ProductCategory{
			ID:       3,
			Category: "Aksesoris Merchandise",
			Icon:     "https://www.riodigitalprint.com/liteprint/application/liteprint/rio/assets/product_category/e5ccd28a1b37fc4523864fcda30149d5.png"}},
		{Sedeer: &domain.ProductCategory{
			ID:       4,
			Category: "Fotak & Frame",
			Icon:     "https://www.riodigitalprint.com/liteprint/application/liteprint/rio/assets/product_category/160311492323-fotak_untuk_web_blue.png"}},
	}
}

func getProductSedeers() []Sedeer {
	productIDs := []string{"uuidp1", "uuidp2", "uuidp3"}

	return []Sedeer{
		// product 1
		{Sedeer: &domain.Product{
			ID:                productIDs[0],
			ProductName:       "Cetak Spanduk",
			ProductCategoryID: 1,
			Price:             25000,
			IsAvailable:       true,
			Description:       "Spanduk dicetak dibahan Flexi China 280 Gsm",
			MinOrder:          1,
			CreatedAt:         time.Now(),
			UpdatedAt:         time.Now(),
		}},
		{Sedeer: &domain.ProductImage{
			ID:        "img1p1",
			Image:     "https://www.riodigitalprint.com/liteprint/application/liteprint/rio/assets/product_image/158469240674-spanduk.jpg",
			ProductID: productIDs[0],
		}},
		{Sedeer: &domain.ProductImage{
			ID:        "img2p1",
			Image:     "https://www.riodigitalprint.com/liteprint/application/liteprint/rio/assets/product_image/158469240674-spanduk-a.jpg",
			ProductID: productIDs[0],
		}},
		// product 2
		{Sedeer: &domain.Product{
			ID:                productIDs[1],
			ProductName:       "Roll Banner",
			ProductCategoryID: 1,
			Price:             250000,
			IsAvailable:       true,
			Description:       "Stand Roll Banner Uk. 60x160cm Bahan roll banner alumunium free tas roll banner",
			MinOrder:          1,
			CreatedAt:         time.Now(),
			UpdatedAt:         time.Now(),
		}},
		{Sedeer: &domain.ProductImage{
			ID:        "img1p2",
			Image:     "https://www.riodigitalprint.com/liteprint/application/liteprint/rio/assets/product_image/1584693200182-ROLLBANNER-KOSONGAN.jpg",
			ProductID: productIDs[1],
		}},
		// product 3
		{Sedeer: &domain.Product{
			ID:                productIDs[2],
			ProductName:       "X-Banner",
			ProductCategoryID: 1,
			Price:             50000,
			IsAvailable:       true,
			Description:       "Kaki X Banner Tanpa Cetak",
			MinOrder:          1,
			CreatedAt:         time.Now(),
			UpdatedAt:         time.Now(),
		}},
		{Sedeer: &domain.ProductImage{
			ID:        "img1p3",
			Image:     "https://www.riodigitalprint.com/liteprint/application/liteprint/rio/assets/product_image/1584692482173-xbanner-kosongan.jpg",
			ProductID: productIDs[2],
		}},
	}
}

func GetOrdersSeeder() []Sedeer {
	return []Sedeer{
		{Sedeer: &domain.Order{
			ID:     "order-1",
			UserID: "114169873760739514656",
			Products: []domain.OrderItem{
				{ID: 1,
					OrderID:    "order-1",
					ProductID:  "uuidp1",
					Quantity:   1,
					TotalPrice: 5000,
				},
				{ID: 2,
					OrderID:    "order-1",
					ProductID:  "uuidp2",
					Quantity:   2,
					TotalPrice: 15000,
				},
			},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}},
	}
}
