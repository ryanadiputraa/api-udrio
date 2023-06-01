package cmd

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ryanadiputraa/api-udrio/pkg/cors"
	"github.com/ryanadiputraa/api-udrio/pkg/database"
	"github.com/spf13/viper"

	_adminDelivery "github.com/ryanadiputraa/api-udrio/app/admin/delivery"
	_adminHandler "github.com/ryanadiputraa/api-udrio/app/admin/handler"
	_adminRepository "github.com/ryanadiputraa/api-udrio/app/admin/repository"

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

	_orderDelivery "github.com/ryanadiputraa/api-udrio/app/order/delivery"
	_orderHandler "github.com/ryanadiputraa/api-udrio/app/order/handler"
	_orderRepository "github.com/ryanadiputraa/api-udrio/app/order/repository"
)

func serveHTTP() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*")
	r.Static("/assets", "./assets")
	r.MaxMultipartMemory = 5 << 20 // max 8 MiB
	r.SetTrustedProxies(nil)

	// Middlewares
	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	r.Use(cors.CORSMiddleware())

	oauth2 := r.Group("/oauth")
	api := r.Group("/api")
	admin := r.Group("/admin")

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
	productRepository := _productRepository.NewProductRepository(database.DB, RedisClient)
	productHandler := _productHandler.NewProductHandler(productRepository)
	_productDelivery.NewProductDelivery(api, productHandler)

	// Orders
	orderRepository := _orderRepository.NewOrderRepository(database.DB)
	orderHandler := _orderHandler.NewOrderHandler(orderRepository)
	_orderDelivery.NewOrderDelivery(api, orderHandler)

	// admin
	adminRepository := _adminRepository.NewAdminRepository(database.DB, RedisClient)
	adminHandler := _adminHandler.NewAdminHandler(adminRepository)
	_adminDelivery.NewAdminDelivery(admin, adminHandler, productHandler)

	// Setup server port & handler
	port := viper.GetString("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	r.Run(fmt.Sprintf(":%s", port))
}
