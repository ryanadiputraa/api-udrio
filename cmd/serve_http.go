package cmd

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ryanadiputraa/api-udrio/pkg/cors"
	"github.com/ryanadiputraa/api-udrio/pkg/database"
	"github.com/spf13/viper"

	_oauthDelivery "github.com/ryanadiputraa/api-udrio/app/oauth/delivery"
	_oauthHandler "github.com/ryanadiputraa/api-udrio/app/oauth/handler"

	_productDelivery "github.com/ryanadiputraa/api-udrio/app/product/delivery"
	_productHandler "github.com/ryanadiputraa/api-udrio/app/product/handler"
	_productRepository "github.com/ryanadiputraa/api-udrio/app/product/repository"
)

func serveHTTP() {

	r := gin.Default()
	r.SetTrustedProxies(nil)

	// Middlewares
	r.Use(cors.CORSMiddleware())

	oauth2 := r.Group("/oauth2")
	api := r.Group("/api")

	// Oauth2
	oAuthHandler := _oauthHandler.NewOAuthHandler()
	_oauthDelivery.NewOAuthDelivery(oauth2, oAuthHandler)

	// Products
	productRepository := _productRepository.NewProductRepository(database.DB)
	productHandler := _productHandler.NewProductHandler(productRepository)
	_productDelivery.NewProductDelivery(api, productHandler)

	// Setup server port & handler
	port := viper.GetString("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	r.Run(fmt.Sprintf(":%s", port))
}
