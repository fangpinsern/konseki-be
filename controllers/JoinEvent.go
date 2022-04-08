package controllers

import (
	"github.com/fangpinsern/konseki-be/logger"
	"github.com/fangpinsern/konseki-be/services"
	"github.com/fangpinsern/konseki-be/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JoinEventController(c *gin.Context){
	userId := util.GetUserId(c)

	var joinEventRequest JoinEventRequest
	var joinEventResponse JoinEventResponse

	if err := c.BindJSON(&joinEventRequest); err != nil {
		joinEventResponse.IsSuccess = false
		c.JSON(http.StatusBadRequest, joinEventResponse)
		return
	}

	eventId := joinEventRequest.Id
	updatedEvent, err := services.JoinEvent(c, eventId, userId)
	if err != nil {
		joinEventResponse.IsSuccess = false
		c.JSON(http.StatusBadRequest, joinEventResponse)
		return
	}

	creatorProfile, err := services.GetProfile(c, updatedEvent.CreatorId)
	if err != nil {
		joinEventResponse.IsSuccess = false
		logger.LogInternal(c, err, "creator profile not found for this event")
		c.JSON(http.StatusBadRequest, joinEventResponse)
		return
	}

	joinEventResponse.IsSuccess = true
	joinEventResponse.EventName = updatedEvent.Name
	joinEventResponse.Id = eventId
	joinEventResponse.CreatorName = creatorProfile.Name

	c.JSON(http.StatusOK, joinEventResponse)
	return
}
