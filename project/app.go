package main

import (
	"net/http"

	"cloudterms.net/project/models"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.GET("/events", getEvents)
	server.POST("/create_event", createEvent)

	server.Run(":8080")

}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	event.UserID = 1
	event.ID = 1

	context.JSON(http.StatusCreated, gin.H{"message": "Event Created!", "event": event})

	// save event
	event.Save()
}
