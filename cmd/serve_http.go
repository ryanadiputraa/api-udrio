package cmd

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ryanadiputraa/api-udrio/pkg/cors"
	"github.com/ryanadiputraa/api-udrio/pkg/database"
	"github.com/spf13/viper"

	_oauthDelivery "github.com/ryanadiputraa/api-udrio/app/oauth/delivery"
	_oauthHandler "github.com/ryanadiputraa/api-udrio/app/oauth/handler"

	_userDelivery "github.com/ryanadiputraa/api-udrio/app/user/delivery"
	_userHandler "github.com/ryanadiputraa/api-udrio/app/user/handler"
	_userRepository "github.com/ryanadiputraa/api-udrio/app/user/repository"

	_cartDelivery "github.com/ryanadiputraa/api-udrio/app/cart/delivery"
	_cartHandler "github.com/ryanadiputraa/api-udrio/app/cart/handler"
	_cartRepository "github.com/ryanadiputraa/api-udrio/app/cart/repository"

	_productDelivery "github.com/ryanadiputraa/api-udrio/app/product/delivery"
	_productHandler "github.com/ryanadiputraa/api-udrio/app/product/handler"
	_productRepository "github.com/ryanadiputraa/api-udrio/app/product/repository"
)

func serveHTTP() {

	r := gin.Default()
	r.SetTrustedProxies(nil)

	// Middlewares
	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	r.Use(cors.CORSMiddleware())

	oauth2 := r.Group("/oauth")
	api := r.Group("/api")

	// cart
	cartRepository := _cartRepository.NewCartRepository(database.DB)
	cartHandler := _cartHandler.NewCartHandler(cartRepository)
	_cartDelivery.NewCartDelivery(api, cartHandler)

	// user
	userRepository := _userRepository.NewUserRepository(database.DB)
	userHandler := _userHandler.NewUserHandler(userRepository, cartRepository)
	_userDelivery.NewUserDelivery(api, AuthMiddleware(), userHandler)

	// Oauth2
	oAuthHandler := _oauthHandler.NewOAuthHandler()
	_oauthDelivery.NewOAuthDelivery(oauth2, oAuthHandler, userHandler)

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
