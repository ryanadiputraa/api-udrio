package cmd

import (
	"log"

	"github.com/spf13/viper"
)

func init() {
	loadConfig()
}

func loadConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName(".udrio")
	viper.SetConfigType("yml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("load config: %s", err.Error())
	}
}

func Execute() {
	serveHTTP()
}
