package handler

import (
	"net/http"
	"learning-platform/internal/service"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	// TODO: Implement registration logic
	c.JSON(http.StatusOK, gin.H{"message": "Register endpoint"})
}

func (h *AuthHandler) Login(c *gin.Context) {
	// TODO: Implement login logic
	c.JSON(http.StatusOK, gin.H{"message": "Login endpoint"})
}

func (h *AuthHandler) GetMe(c *gin.Context) {
	// TODO: Implement get current user logic
	c.JSON(http.StatusOK, gin.H{"message": "GetMe endpoint"})
}