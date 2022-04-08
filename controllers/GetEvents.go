package controllers

import (
	"errors"
	"github.com/fangpinsern/konseki-be/logger"
	"github.com/fangpinsern/konseki-be/services"
	"github.com/fangpinsern/konseki-be/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetEventsController(c *gin.Context){
	userId := util.GetUserId(c)

	if userId == "" {
		err := errors.New("please login")
		logger.LogInternal(c, err, "Profile Data not found")
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	responseEventList, err := services.GetAttendedEvents(c, userId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var getEventsResponse GetEventsResponse
	getEventsResponse.Events = responseEventList

	c.IndentedJSON(http.StatusAccepted, getEventsResponse)
}
