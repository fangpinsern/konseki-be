package controllers

import (
	"github.com/gin-gonic/gin"
	"konseki-be/db"
	"konseki-be/logger"
	"konseki-be/services"
	"konseki-be/util"
	"net/http"
	"time"
)

// Controller creates an event and stores it in the database
// The creator is tagged to the user

func CreateEventController(c *gin.Context) {
	var createEventRequest CreateEventRequest
	var createEventResponse CreateEventResponse

	if err := c.BindJSON(&createEventRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	eventName := createEventRequest.Name
	eventDate := time.Now().Unix()

	creatorId := util.GetUserId(c)

	event := services.Event{ //use uuid
		Name:      eventName,
		Attended:  []string{creatorId},
		Date:      eventDate,
		CreatorId: creatorId,
	}

	//result, err := db.DatabaseClient.Collection("event").Doc(""). Set(c, event)
	result,_, err :=db.EventCollection.Add(c, event)
	if err != nil {
		logger.LogInternal(c, err, "Something went wrong")
	}

	logger.LogInternal(c, nil, result.ID)

	joinEventLink := "https://google.com/"
	eventId := result.ID
	createEventResponse.Link = joinEventLink + "?eventId=" + eventId
	createEventResponse.Name = eventName
	createEventResponse.Id = eventId


	c.IndentedJSON(http.StatusAccepted, createEventResponse)

}
