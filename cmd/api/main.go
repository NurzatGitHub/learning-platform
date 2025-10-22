package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Basic health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	})

	// Public routes
	router.POST("/register", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Register endpoint"})
	})
	router.POST("/login", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Login endpoint"})
	})

	// Protected routes group
	api := router.Group("/api")
	{
		api.GET("/users/me", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "GetMe endpoint"})
		})
	}

	log.Println("Starting server on :8081")
	if err := router.Run(":8081"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}