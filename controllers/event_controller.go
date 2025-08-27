package controllers

import (
	"net/http"
	"strconv"
	"ticketing-go/models"
	"ticketing-go/services"

	"github.com/gin-gonic/gin"
)

type EventController struct {
	service services.EventService
}

func NewEventController(service services.EventService) *EventController {
	return &EventController{service}
}

func (c *EventController) CreateEvent(ctx *gin.Context) {
	var event models.SubmitEvent

	userID := ctx.GetUint("user_id")
	event.CreatedBy = uint64(userID)

	if err := ctx.ShouldBindJSON(&event); err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:        "error",
			Message:       "Invalid request body",
			MessageDetail: err.Error(),
		})
		return
	}

	if err := c.service.CreateEvent(&event); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status:        "error",
			Message:       "Failed to create event",
			MessageDetail: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, models.Response{
		Status:  "success",
		Message: "Event created successfully",
		Data:    event,
	})
}

func (c *EventController) GetEvents(ctx *gin.Context) {
	events, err := c.service.GetEvents()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status:        "error",
			Message:       "Failed to retrieve events",
			MessageDetail: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Status:  "success",
		Message: "Events retrieved successfully",
		Data:    events,
	})
}

func (c *EventController) DeleteEvent(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := c.service.DeleteEvent(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status:        "error",
			Message:       "Failed to delete event",
			MessageDetail: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Status:  "success",
		Message: "Event deleted successfully",
	})
}

func (c *EventController) UpdateEvent(ctx *gin.Context) {
	var event models.UpdateEvent
	id := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&event); err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:        "error",
			Message:       "Invalid request body",
			MessageDetail: err.Error(),
		})
		return
	}

	// Convert id from string to uint64
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			Status:        "error",
			Message:       "Invalid event ID",
			MessageDetail: err.Error(),
		})
		return
	}
	event.ID = idUint

	if err := c.service.UpdateEvent(&event); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status:        "error",
			Message:       "Failed to update event",
			MessageDetail: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Status:  "success",
		Message: "Event updated successfully",
		Data:    event,
	})
}

	
