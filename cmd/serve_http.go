package cmd

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ryanadiputraa/api-udrio/app/product/handler"
	"github.com/ryanadiputraa/api-udrio/app/product/repository"
	"github.com/ryanadiputraa/api-udrio/app/product/service"
	"github.com/ryanadiputraa/api-udrio/pkg/database"
	"github.com/spf13/viper"
)

func serveHTTP() {

	r := gin.Default()
	r.SetTrustedProxies(nil)
	api := r.Group("/api")

	// Products
	productRepository := repository.NewProductRepository(database.DB)
	productService := service.NewProductService(productRepository)
	handler.NewProductHandler(api, productService)

	// Setup server port & handler
	port := viper.GetString("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	r.Run(fmt.Sprintf(":%s", port))
}
