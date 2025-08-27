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

		userId, role, err := utils.ValidateToken(parts[1])
		if err != nil {
			c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.JSON(401, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		c.Set("user_id", userId)
		c.Set("user_role", role)
		c.Next()
	}
}



func RoleAdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		roleInterface, exists := c.Get("user_role")

		if !exists || roleInterface != "Admin" {
			c.JSON(403, gin.H{"error": "Only admin can access this resource"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func RoleUserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("user_role")
		if !exists || userRole != "User" {
			c.JSON(403, gin.H{"error": "Only user can access this resource"})
			c.Abort()
			return
		}

		c.Next()
	}
}