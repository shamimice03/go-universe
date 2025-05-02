package routes

import (
	"log"
	"net/http"
	"strconv"

	"cloudterms.net/project/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	// userId set by middleware
	userId := context.GetInt64("userID")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the event"})
		return
	}

	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to register user for the event."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Successfully registered for the event!"})

}

func cancelRegistration(context *gin.Context) {
	// userId set by middleware
	userId := context.GetInt64("userID")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	var event models.Event
	event.ID = eventId

	err = event.CancelRegistration(userId)
	if err != nil {
		log.Printf("Event Cancel registration failed for %v", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to cancel user for the event."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User registration cancelled!"})
}
