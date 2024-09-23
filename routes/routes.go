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
		bike.PUT("/:id", middleware.Authentication("user"), controllers.UpdateBike)
		bike.DELETE("/:id", middleware.Authentication("user"), controllers.DeleteBikeById)
		bike.GET("/:id", controllers.GetBikeDetail)
	}

	r.Run(os.Getenv("PORT"))
}
