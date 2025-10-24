package middleware

import (
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: Implement JWT authentication
		// For now, just pass through
		c.Next()
	}
}