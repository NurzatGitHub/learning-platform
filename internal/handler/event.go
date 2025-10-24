package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type EventHandler struct{}

func NewEventHandler() *EventHandler {
	return &EventHandler{}
}

func (h *EventHandler) CreateEvent(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Create event endpoint"})
}

func (h *EventHandler) GetEvents(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get all events endpoint"})
}

func (h *EventHandler) GetEvent(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get event by ID endpoint"})
}

func (h *EventHandler) UpdateEvent(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Update event endpoint"})
}

func (h *EventHandler) DeleteEvent(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Delete event endpoint"})
}