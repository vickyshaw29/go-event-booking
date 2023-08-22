package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vickyshaw29/events/database"
	"github.com/vickyshaw29/events/routes"
)

func main() {
	r := gin.Default()
	database.InitDB()
	routes.RegisterRoutes(r)
	r.Run(":8080")
}
