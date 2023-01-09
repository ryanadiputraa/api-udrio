package cmd

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ryanadiputraa/api-udrio/app/product/handler"
	"github.com/ryanadiputraa/api-udrio/pkg/database"
	"github.com/spf13/viper"
)

func serveHTTP() {
	_ = database.GetConnection()

	r := gin.Default()
	api := r.Group("/api")
	handler.NewProductHandler(api)

	// Setup server port & handler
	port := viper.GetString("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	r.Run(fmt.Sprintf(":%s", port))
}
