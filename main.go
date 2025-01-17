package main

import (
	"example.com/go-rest-api/db"
	"example.com/go-rest-api/middlewares"
	"example.com/go-rest-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	server.Use(middlewares.Cors())
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
