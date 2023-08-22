package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vickyshaw29/events/handlers"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/events", handlers.GetAllEvents)
}
