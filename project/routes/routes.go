package routes

import (
	"cloudterms.net/project/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	// One option
	// server.POST("/create_event", middlewares.Authenticate, createEvent)
	// server.PUT("/update_event/:id", middlewares.Authenticate, updateEvent)
	// server.DELETE("/events/:id", middlewares.Authenticate, deleteEvent)

	// Other Option
	authenticate := server.Group("/")
	authenticate.Use(middlewares.Authenticate)
	authenticate.POST("/create_event", createEvent)
	authenticate.PUT("/update_event/:id", updateEvent)
	authenticate.DELETE("/events/:id", deleteEvent)
	authenticate.POST("/events/:id/register", registerForEvent)
	authenticate.DELETE("/events/:id/cancel_registration", cancelRegistration)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
