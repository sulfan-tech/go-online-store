package router

import (
	cartService "go-online-store/internal/domain/cart/service"
	customerService "go-online-store/internal/domain/customer/service"
	orderService "go-online-store/internal/domain/order/service"
	productService "go-online-store/internal/domain/product/service"
	"go-online-store/internal/handlers/cart"
	"go-online-store/internal/handlers/customer"
	"go-online-store/internal/handlers/order"
	"go-online-store/internal/handlers/product"
	"go-online-store/internal/middleware/jwt"
	"go-online-store/pkg/logger"

	_ "go-online-store/server/cmd/docs"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func RegisterRouter(e *echo.Echo, log *logger.Logger) *echo.Echo {
	// e.Use(jwt.ValidateJWT)

	// Init Service
	userService := customerService.NewInstanceUserService()
	productService := productService.NewInstanceProductService()
	cartService := cartService.NewInstanceCartService()
	orderService, _ := orderService.NewOrderService()

	// Init Handler
	customerHandler := customer.NewCustomerHandler(userService)
	productHandler := product.NewProductHandler(productService)
	cartHandler := cart.NewCartHandler(cartService)
	orderHandler := order.NewOrderHandler(orderService)

	// Group routes for API v1
	v1 := e.Group("/v1")

	// Routes for customer authentication
	v1.POST("/user/login", customerHandler.CustomerLogin)
	v1.POST("/user/register", customerHandler.CustomerRegister)

	// Router for product
	v1.GET("/products", jwt.ValidateJWT(productHandler.GetProductsByCategoryHandler))

	// Routes for cart
	v1.GET("/cart", jwt.ValidateJWT(cartHandler.GetCartHandler))
	v1.POST("/cart", jwt.ValidateJWT(cartHandler.AddToCartHandler))
	v1.DELETE("/cart", jwt.ValidateJWT(cartHandler.RemoveFromCartHandler))

	// Routes for order
	v1.POST("/checkout", jwt.ValidateJWT(orderHandler.CheckoutHandler))
	v1.POST("/checkout/paid", jwt.ValidateJWT(orderHandler.TransactionPaidHandler))

	// Swagger endpoint
	v1.GET("/swagger/*", echoSwagger.WrapHandler)

	return e
}
