package database

import (
	"fmt"

	"github.com/ryanadiputraa/api-udrio/domain"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GetConnection() {
	dsn := viper.GetString("POSTGRES_DSN")
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("db connection: %s", err))
	}

	db, err := conn.DB()
	if err != nil {
		panic(fmt.Sprintf("db connection: %s", err))
	}
	db.SetMaxIdleConns(viper.GetInt("MAX_IDLE"))
	db.SetMaxOpenConns(viper.GetInt("MAX_CONNS"))
	db.SetConnMaxIdleTime(viper.GetDuration("IDLE_TIME"))
	db.SetConnMaxLifetime(viper.GetDuration("LIFE_TIME"))

	DB = conn

	makeMigration()
}

func makeMigration() {
	DB.AutoMigrate(
		&domain.Admin{},
		&domain.Cart{},
		&domain.CartItem{},
		&domain.User{},
		&domain.Product{},
		&domain.ProductImage{},
		&domain.ProductCategory{},
		&domain.Order{},
		&domain.OrderItem{},
	)
}
