package routes

import (
	"ticketing-go/controllers"
	"ticketing-go/repositories"
	"ticketing-go/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupUserRoutes(router *gin.RouterGroup, db *gorm.DB) {
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	protected := router.Group("/")
	{
		protected.POST("/register", userController.Register)
		protected.POST("/login", userController.Login)
	}
}