package database

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetConnection() *gorm.DB {
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

	return conn
}
