package controllers

import (
	"github.com/gin-gonic/gin"
	"konseki-be/services"
	"konseki-be/util"
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

	joinEventResponse.IsSuccess = true
	joinEventResponse.EventName = updatedEvent.Name
	joinEventResponse.Id = eventId
	c.JSON(http.StatusOK, joinEventResponse)
	return
}
