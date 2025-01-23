package routes

import (
	"example.com/go-rest-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	const eventPath string = "/conferences"

	// ROUTES EVENT
	server.GET(eventPath, getEvents)
	server.GET(eventPath+"/:id", getEvent)

	//ROUTES AUTH
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST(eventPath, createEvent)
	authenticated.PUT(eventPath+"/:id", updateEvent)
	authenticated.DELETE(eventPath+"/:id", deleteEvent)
	authenticated.POST(eventPath+"/:id/register", registerForEvent)
	authenticated.DELETE(eventPath+"/:id/register", cancelRegistration)

	// ROUTES USER
	server.POST("/signup", signup)
	server.POST("/login", login)
}
