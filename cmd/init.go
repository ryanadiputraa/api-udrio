package cmd

import (
	"log"

	"github.com/ryanadiputraa/api-udrio/migrations"
	"github.com/spf13/viper"
)

func init() {
	loadConfig()
	migrations.Migrate()
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
	serveHTTP()
}
