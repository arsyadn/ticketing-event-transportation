package routes

import (
	"ticketing-go/controllers"
	"ticketing-go/middleware"
	"ticketing-go/repositories"
	"ticketing-go/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupTicketRoutes(router *gin.RouterGroup, db *gorm.DB) {
	TicketRepo := repositories.NewTicketRepository(db)
	EventRepo := repositories.NewEventRepository(db)
	UserRepo := repositories.NewUserRepository(db)
	TicketService := services.NewTicketService(TicketRepo, EventRepo, UserRepo)
	TicketController := controllers.NewTicketController(TicketService)

	// admin := router.Group("/tickets")
	// admin.Use(middleware.AuthMiddleware())
	// admin.Use(middleware.RoleAdminMiddleware())

	protected := router.Group("/tickets")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("", TicketController.CreateTicket)
		protected.GET("", TicketController.GetAllTickets)
		protected.GET("/:id", TicketController.GetTicketByID)
	}
}