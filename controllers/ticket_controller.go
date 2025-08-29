package controllers

import (
	"fmt"
	"strconv"
	"ticketing-go/models"
	"ticketing-go/services"

	"github.com/gin-gonic/gin"
)

type TicketController struct {
	service services.TicketService
}

func NewTicketController(service services.TicketService) *TicketController {
	return &TicketController{
		service: service,
	}
}

func (c *TicketController) CreateTicket(ctx *gin.Context) {
	var ticket models.BuyTicket

	userID := ctx.GetUint("user_id")
	ticket.UserID = uint64(userID)

	if err := ctx.ShouldBindJSON(&ticket); err != nil {
		ctx.JSON(400, models.ErrorResponse{
			Status:        "error",
			Message:       "Invalid request body",
			MessageDetail: err.Error(),
		})
		return
	}

	if err := c.service.CreateTicket(&ticket); err != nil {
		ctx.JSON(500, models.ErrorResponse{
			Status:        "error",
			Message:       "Failed to create ticket",
			MessageDetail: err.Error(),
		})
		return
	}

	ctx.JSON(200, models.Response{
		Status:  "success",
		Message: "Ticket created successfully",
		Data:   ticket,
	})
}

func (c *TicketController) GetAllTickets(ctx *gin.Context) {
	
	userRole, _ := ctx.Get("user_role")
	userID := ctx.GetUint("user_id")

	tickets, err := c.service.GetAllTickets(uint64(userID), fmt.Sprintf("%v", userRole))

	if err != nil {
		ctx.JSON(500, models.ErrorResponse{
			Status:        "error",
			Message:       "Failed to retrieve tickets",
			MessageDetail: err.Error(),
		})
		return
	}

	ctx.JSON(200, models.Response{
		Status:  "success",
		Message: "Tickets retrieved successfully",
		Data:   tickets,
	})
}

func (c *TicketController) GetTicketByID(ctx *gin.Context) {
	ticketIDStr := ctx.Param("id")
	ticketID, err := strconv.ParseUint(ticketIDStr, 10, 64)

	if err != nil {
		ctx.JSON(400, models.ErrorResponse{
			Status:        "error",
			Message:       "Invalid ticket ID",
			MessageDetail: err.Error(),
		})
		return
	}
	
	userRole, _ := ctx.Get("user_role")
	userID := ctx.GetUint("user_id")

	ticket, err := c.service.GetTicketByID(ticketID, uint64(userID), fmt.Sprintf("%v", userRole))
	if err != nil {
		ctx.JSON(500, models.ErrorResponse{
			Status:        "error",
			Message:       "Failed to retrieve ticket",
			MessageDetail: err.Error(),
		})
		return
	}

	ctx.JSON(200, models.Response{
		Status:  "success",
		Message: "Ticket retrieved successfully",
		Data:   ticket,
	})
}