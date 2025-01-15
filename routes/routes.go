package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	const eventPath string = "/events"

	// ROUTES EVENT
	server.GET(eventPath, getEvents)
	server.GET(eventPath+"/:id", getEvent)
	server.POST(eventPath, createEvent)
	server.PUT(eventPath+"/:id", updateEvent)
	server.DELETE(eventPath+"/:id", deleteEvent)

	// ROUTES USER
	server.POST("/signup", signup)
	server.POST("/login", login)
}
