package routes

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/markaz/controllers"
)

// SetupRouter ... Configure routes for the service
func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	apis := r.Group("/api")
	orders := apis.Group("/orders")
	{
		orders.POST("create_order", controllers.CreateOrder)
		orders.GET("customer_orders", controllers.GetCustomerOrders)
	}
	return r
}
