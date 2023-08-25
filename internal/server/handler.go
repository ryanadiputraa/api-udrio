package server

import (
	_adminDelivery "github.com/ryanadiputraa/api-udrio/app/admin/delivery"
	_adminHandler "github.com/ryanadiputraa/api-udrio/app/admin/handler"
	_adminRepository "github.com/ryanadiputraa/api-udrio/app/admin/repository"
	"github.com/ryanadiputraa/api-udrio/internal/middleware"

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

func (s *Server) MapHandlers() {
	oauth2 := s.gin.Group("/oauth")
	api := s.gin.Group("/api")
	admin := s.gin.Group("/admin")

	// cart
	cartRepository := _cartRepository.NewCartRepository(s.db)
	cartHandler := _cartHandler.NewCartHandler(cartRepository)
	_cartDelivery.NewCartDelivery(api, cartHandler)

	// user
	userRepository := _userRepository.NewUserRepository(s.db)
	userHandler := _userHandler.NewUserHandler(userRepository, cartRepository)
	_userDelivery.NewUserDelivery(api, middleware.AuthMiddleware(), userHandler)

	// Oauth2
	oAuthHandler := _oauthHandler.NewOAuthHandler()
	_oauthDelivery.NewOAuthDelivery(oauth2, oAuthHandler, userHandler)

	// Products
	productRepository := _productRepository.NewProductRepository(s.db, s.redis)
	productHandler := _productHandler.NewProductHandler(productRepository)
	_productDelivery.NewProductDelivery(api, productHandler)

	// Orders
	orderRepository := _orderRepository.NewOrderRepository(s.db)
	orderHandler := _orderHandler.NewOrderHandler(orderRepository)
	_orderDelivery.NewOrderDelivery(api, orderHandler)

	// admin
	adminRepository := _adminRepository.NewAdminRepository(s.db, s.redis)
	adminHandler := _adminHandler.NewAdminHandler(adminRepository)
	_adminDelivery.NewAdminDelivery(admin, adminHandler, productHandler)
}
