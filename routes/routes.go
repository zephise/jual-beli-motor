package routes

import (
	"jual-beli-motor/controllers"
	"jual-beli-motor/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

func Routes() {
	r := gin.Default()
	r.GET("/health-check", controllers.HealthCheck)
	r.POST("/auth/login", controllers.Login)
	r.POST("/auth/register", controllers.CreateUserNonAdmin)
	r.POST("/auth/admin/register", controllers.CreateUserAdmin)

	bikeType := r.Group("/bike-type")
	{
		bikeType.POST("/", controllers.CreatedBikeType)
		bikeType.GET("/", controllers.GetAllBikeType)
		bikeType.GET("/:id", controllers.GetBikeTypeById)
		bikeType.PUT("/:id", controllers.UpdateBikeType)
		bikeType.DELETE("/:id", controllers.DeleteBikeTypeById)
	}

	bike := r.Group("/bike")
	{
		bike.POST("/", middleware.Authentication("user"), controllers.CreateBike)
		bike.GET("/", middleware.Authentication("user"), controllers.GetAllBike)
		bike.GET("/check-bike-status", middleware.Authentication("user"), controllers.CheckBikeStatus)
		bike.PUT("/:id", middleware.Authentication("user"), controllers.UpdateBike)
		bike.DELETE("/:id", middleware.Authentication("user"), controllers.DeleteBikeById)
		bike.GET("/:id", controllers.GetBikeDetail)
		bike.POST("/purchase/:id", middleware.Authentication("user"), controllers.PurchaseBike)
	}

	coupon := r.Group("/coupon")
	{
		coupon.POST("/", middleware.Authentication("admin"), controllers.CreateCoupon)

	}

	r.Run(os.Getenv("PORT"))
}
