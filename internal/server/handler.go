package server

import (
	"github.com/ryanadiputraa/api-udrio/internal/middleware"

	_adminDelivery "github.com/ryanadiputraa/api-udrio/app/admin/delivery"
	_adminHandler "github.com/ryanadiputraa/api-udrio/app/admin/handler"
	_adminRepository "github.com/ryanadiputraa/api-udrio/app/admin/repository"

	_oauthDelivery "github.com/ryanadiputraa/api-udrio/internal/oauth/delivery"
	_oauthUsecase "github.com/ryanadiputraa/api-udrio/internal/oauth/usecase"

	_userDelivery "github.com/ryanadiputraa/api-udrio/internal/user/delivery"
	_userRepository "github.com/ryanadiputraa/api-udrio/internal/user/repository"
	_userUsecase "github.com/ryanadiputraa/api-udrio/internal/user/usecase"

	_cartDelivery "github.com/ryanadiputraa/api-udrio/internal/cart/delivery"
	_cartRepository "github.com/ryanadiputraa/api-udrio/internal/cart/repository"
	_cartUsecase "github.com/ryanadiputraa/api-udrio/internal/cart/usecase"

	_productDelivery "github.com/ryanadiputraa/api-udrio/internal/product/delivery"
	_productRepository "github.com/ryanadiputraa/api-udrio/internal/product/repository"
	_productUsecase "github.com/ryanadiputraa/api-udrio/internal/product/usecase"

	_orderDelivery "github.com/ryanadiputraa/api-udrio/internal/order/delivery"
	_orderRepository "github.com/ryanadiputraa/api-udrio/internal/order/repository"
	_orderUsecase "github.com/ryanadiputraa/api-udrio/internal/order/usecase"
)

func (s *Server) MapHandlers() {
	oauth2 := s.gin.Group("/oauth")
	api := s.gin.Group("/api")
	admin := s.gin.Group("/admin")

	// cart
	cartRepository := _cartRepository.NewCartRepository(s.db)
	cartUsecase := _cartUsecase.NewCartUsecase(cartRepository)
	_cartDelivery.NewCartDelivery(api, cartUsecase)

	// user
	userRepository := _userRepository.NewUserRepository(s.db)
	userUsecase := _userUsecase.NewUserUsecase(userRepository, cartRepository)
	_userDelivery.NewUserDelivery(api, middleware.AuthMiddleware(), userUsecase)

	// Oauth2
	oauthUsecase := _oauthUsecase.NewOAuthUsecase()
	_oauthDelivery.NewOAuthDelivery(oauth2, *s.conf, oauthUsecase, userUsecase)

	// Products
	productRepository := _productRepository.NewProductRepository(s.db, s.redis, s.storage)
	productUsecase := _productUsecase.NewProductUsecase(productRepository)
	_productDelivery.NewProductDelivery(api, productUsecase)

	// Orders
	orderRepository := _orderRepository.NewOrderRepository(s.db)
	orderHandler := _orderUsecase.NewOrderUsecase(orderRepository)
	_orderDelivery.NewOrderDelivery(api, orderHandler)

	// admin
	adminRepository := _adminRepository.NewAdminRepository(s.db, s.redis)
	adminHandler := _adminHandler.NewAdminHandler(adminRepository)
	_adminDelivery.NewAdminDelivery(admin, adminHandler, productUsecase)
}
