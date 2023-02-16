package cmd

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ryanadiputraa/api-udrio/pkg/cors"
	"github.com/ryanadiputraa/api-udrio/pkg/database"
	"github.com/spf13/viper"

	_oauthHandler "github.com/ryanadiputraa/api-udrio/app/oauth/handler"
	_oauthService "github.com/ryanadiputraa/api-udrio/app/oauth/service"

	_productHandler "github.com/ryanadiputraa/api-udrio/app/product/handler"
	_productRepository "github.com/ryanadiputraa/api-udrio/app/product/repository"
	_productService "github.com/ryanadiputraa/api-udrio/app/product/service"
)

func serveHTTP() {

	r := gin.Default()
	r.SetTrustedProxies(nil)

	// Middlewares
	r.Use(cors.CORSMiddleware())

	oauth2 := r.Group("/oauth2")
	api := r.Group("/api")

	// Oauth2
	oauthService := _oauthService.NewOAuthService()
	_oauthHandler.NewOauthHandler(oauth2, oauthService)

	// Products
	productRepository := _productRepository.NewProductRepository(database.DB)
	productService := _productService.NewProductService(productRepository)
	_productHandler.NewProductHandler(api, productService)

	// Setup server port & handler
	port := viper.GetString("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	r.Run(fmt.Sprintf(":%s", port))
}
