package cmd

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ryanadiputraa/api-udrio/pkg/cache"
	"github.com/ryanadiputraa/api-udrio/pkg/database"
	"github.com/spf13/viper"
)

var RedisClient cache.Redis

func init() {
	loadConfig()
	database.GetConnection()
	RedisClient = cache.InitRedis()

	if viper.GetString("ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
}

func loadConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("load config: %s", err.Error())
	}
}

func Execute() {
	env := viper.GetString("ENV")
	if env == "dev" {
		err := database.DBSeed(database.DB)
		if err != nil {
			log.Fatalf("db seed: %s", err.Error())
		}
	}
	serveHTTP()
}
