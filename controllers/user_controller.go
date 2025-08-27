package controllers

import (
	"net/http"
	"ticketing-go/models"
	"ticketing-go/services"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	UserService *services.UserService
}

func NewUserController(service *services.UserService) *AuthController {
	return &AuthController{
		UserService: service,
	}
}

func (ac *AuthController) Register(c *gin.Context){
	var user models.User

	if err := c.ShouldBindBodyWithJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	token, userID, err := ac.UserService.Register(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"token": token,
		"user_id": userID,
	})
}

func (ac *AuthController) Login(c *gin.Context) {
	var loginRequest models.LoginRequest

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	token, userID, err := ac.UserService.Login(&loginRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token": token,
		"user_id": userID,
	})
}