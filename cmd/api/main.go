package main

import (
	"log"
	"net/http"
	"os"

	"learning-platform/internal/handler"
	"learning-platform/internal/middleware"
	"learning-platform/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	router := gin.Default()

	// Add middleware
	router.Use(middleware.LoggerMiddleware())

	// Initialize services
	authService := service.NewAuthService("your-jwt-secret")
	
	// Initialize handlers
	authHandler := handler.NewAuthHandler(authService)
	eventHandler := handler.NewEventHandler()
	userHandler := handler.NewUserHandler()

	// Public routes
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "OK", "timestamp": "2025-10-22"})
	})
	router.POST("/register", authHandler.Register)
	router.POST("/login", authHandler.Login)

	// Protected API routes
	api := router.Group("/api")
	api.Use(middleware.AuthMiddleware("your-jwt-secret"))
	{
		// User routes
		api.GET("/users/me", authHandler.GetMe)
		api.GET("/users/profile", userHandler.GetUserProfile)
		api.PUT("/users/profile", userHandler.UpdateUserProfile)

		// Event routes
		api.GET("/events", eventHandler.GetEvents)
		api.GET("/events/:id", eventHandler.GetEvent)
		api.POST("/events", eventHandler.CreateEvent)
		api.PUT("/events/:id", eventHandler.UpdateEvent)
		api.DELETE("/events/:id", eventHandler.DeleteEvent)
	}

	log.Printf("Starting server on :%s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}