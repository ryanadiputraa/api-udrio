package main

import (
	"log"

	"github.com/ryanadiputraa/api-udrio/config"
	"github.com/ryanadiputraa/api-udrio/internal/server"
)

func main() {
	conf, err := config.LoadConfig("config/config")
	if err != nil {
		log.Fatal(err)
	}

	s := server.NewServer(conf)
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
