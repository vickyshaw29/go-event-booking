package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vickyshaw29/events/models"
	"github.com/vickyshaw29/events/utils"
)

func GetEvents(c *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not fetch events try again later.",
		})
		return
	}
	c.JSON(http.StatusOK, events)
}

func GetEventById(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "could not parse event id",
		})
		return
	}
	event, err := models.GetEventById(uint64(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "could not fetch event",
		})
		return
	}
	c.JSON(http.StatusCreated, event)
}

func CreateEvent(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}
	err := utils.VerifyToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
	}
	event := models.Event{}
	err = c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusCreated, gin.H{
			"message": "Could not parse request data",
		})
	}
	event.ID = 1
	event.UserID = 1
	err = event.CreateEvent()
	if err != nil {
		c.JSON(http.StatusCreated, gin.H{
			"message": "Could not create event",
		})
	}
	c.JSON(http.StatusCreated, event)
}

func UpdateEvent(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "could not parse event id",
		})
		return
	}
	_, err = models.GetEventById(uint64(eventId))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "could not fetch event",
		})
		return
	}
	updatedEvent := models.Event{}
	err = c.ShouldBindJSON(&updatedEvent)
	if err != nil {
		c.JSON(http.StatusCreated, gin.H{
			"message": "Could not parse request data",
		})
	}

	updatedEvent.ID = eventId
	err = updatedEvent.UpdateEventById()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "could not update event",
		})
		return
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "Event updated successfully",
	})

}

func DelteEvent(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "could not parse event id",
		})
		return
	}
	event, err := models.GetEventById(uint64(eventId))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "could not fetch event",
		})
		return
	}

	err = event.DeleteEventById()
	if err != nil {
		c.JSON(http.StatusInsufficientStorage, gin.H{
			"message": "Could not delete event",
		})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully!"})

}
