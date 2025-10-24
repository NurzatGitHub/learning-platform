package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) GetUserProfile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get user profile endpoint"})
}

func (h *UserHandler) UpdateUserProfile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Update user profile endpoint"})
}