package controllers

import (
	"github.com/fangpinsern/konseki-be/config"
	"github.com/fangpinsern/konseki-be/db"
	"github.com/fangpinsern/konseki-be/logger"
	"github.com/fangpinsern/konseki-be/services"
	"github.com/fangpinsern/konseki-be/util"
	"github.com/gin-gonic/gin"
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

	joinEventLink := config.GetKonsekiLink()
	eventId := result.ID
	createEventResponse.Link = joinEventLink + "?eventId=" + eventId
	createEventResponse.Name = eventName
	createEventResponse.Id = eventId


	c.IndentedJSON(http.StatusAccepted, createEventResponse)

}
