package main

import (
	"log"

	"github.com/ryanadiputraa/api-udrio/config"
	"github.com/ryanadiputraa/api-udrio/internal/server"
	"github.com/ryanadiputraa/api-udrio/pkg/database"
)

func main() {
	conf, err := config.LoadConfig("config/config")
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.New(conf.Postgres)
	if err != nil {
		log.Fatal(err)
	}

	redis, err := database.InitRedis(conf.Redis.DSN)
	if err != nil {
		log.Fatal(err)
	}

	s := server.NewServer(conf, db, redis)
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
