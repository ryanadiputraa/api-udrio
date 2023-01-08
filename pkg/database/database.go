package database

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetConnection() *gorm.DB {
	dsn := viper.GetString("database.dsn")
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("db connection: %s", err))
	}

	db, err := conn.DB()
	if err != nil {
		panic(fmt.Sprintf("db connection: %s", err))
	}
	db.SetMaxIdleConns(viper.GetInt("database.max_idle"))
	db.SetMaxOpenConns(viper.GetInt("database.max_conns"))
	db.SetConnMaxIdleTime(viper.GetDuration("database.idle_time"))
	db.SetConnMaxLifetime(viper.GetDuration("database.life_time"))

	return conn
}
