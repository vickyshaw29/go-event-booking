package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vickyshaw29/events/handlers"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/events", handlers.GetEvents)
	r.GET("/events/:id", handlers.GetEventById)
	r.POST("/events", handlers.CreateEvent)
	r.PUT("/events/:id", handlers.UpdateEvent)
	r.DELETE("/events/:id", handlers.DelteEvent)

	r.POST("/signup", handlers.Signup)
	r.POST("/login", handlers.Login)
}
