package routes

import (
	"ticketing-go/controllers"
	"ticketing-go/middleware"
	"ticketing-go/repositories"
	"ticketing-go/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupEventRoutes(router *gin.RouterGroup, db *gorm.DB) {
	EventRepo := repositories.NewEventRepository(db)
	EventService := services.NewEventService(EventRepo)
	EventController := controllers.NewEventController(EventService)

	admin := router.Group("/events")
	admin.Use(middleware.AuthMiddleware())
	admin.Use(middleware.RoleAdminMiddleware())

	protected := router.Group("/events")
	protected.Use(middleware.AuthMiddleware())
	{
		admin.POST("", EventController.CreateEvent)
		protected.GET("", EventController.GetEvents)
		admin.DELETE("/:id", EventController.DeleteEvent)
		admin.PUT("/:id", EventController.UpdateEvent)
	}
}