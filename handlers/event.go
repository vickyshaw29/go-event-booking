package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllEvents(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "List of events",
	})
}

func CreateEvent(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"message": "Event created",
	})
}
