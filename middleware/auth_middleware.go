package middleware

import (
	"strings"
	"ticketing-go/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "OPTIONS" {
			c.Next()
			return
		}

		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.JSON(401, gin.H{"error": "Authorization header token is required"})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.JSON(401, gin.H{"error": "Authorization header format must be Bearer token"})
			c.Abort()
			return
		}

		userId, err := utils.ValidateToken(parts[1])
		if err != nil {
			c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.JSON(401, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		c.Set("user_id", userId)
		c.Next()
	}
}