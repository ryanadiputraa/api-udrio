package database

import (
	"github.com/ryanadiputraa/api-udrio/config"
	"github.com/ryanadiputraa/api-udrio/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(conf config.Postgres) (*gorm.DB, error) {
	conn, err := gorm.Open(postgres.Open(conf.DSN), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db, err := conn.DB()
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(conf.MaxIdle)
	db.SetMaxOpenConns(conf.MaxConns)
	db.SetConnMaxIdleTime(conf.IdleTime)
	db.SetConnMaxLifetime(conf.LifeTime)

	makeMigration(conn)

	return conn, nil
}

func makeMigration(db *gorm.DB) {
	db.AutoMigrate(
		&domain.Admin{},
		&domain.Cart{},
		&domain.CartItem{},
		&domain.User{},
		&domain.Product{},
		&domain.ProductImage{},
		&domain.ProductCategory{},
		&domain.Order{},
		&domain.OrderItem{},
		&domain.AssetsPath{},
	)
}
