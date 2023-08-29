package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vickyshaw29/events/handlers"
	"github.com/vickyshaw29/events/middlewares"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/events", handlers.GetEvents)
	r.GET("/events/:id", handlers.GetEventById)
	authenticated := r.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", handlers.CreateEvent)
	authenticated.PUT("/events/:id", handlers.UpdateEvent)
	authenticated.DELETE("/events/:id", handlers.DelteEvent)

	r.POST("/signup", handlers.Signup)
	r.POST("/login", handlers.Login)
}
