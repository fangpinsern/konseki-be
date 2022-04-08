package controllers

import (
	"github.com/fangpinsern/konseki-be/db"
	"github.com/fangpinsern/konseki-be/logger"
	"github.com/fangpinsern/konseki-be/services"
	"github.com/fangpinsern/konseki-be/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterController(c *gin.Context) {
	var registerRequest RegisterRequest
	var registerResponse RegisterResponse

	if err := c.BindJSON(&registerRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	name := registerRequest.Name
	email := util.GetUserEmail(c)
	id := util.GetUserId(c)

	// TODO: save user profile to firebase store
	profile := services.Profile{
		Id:         id,
		Name:       name,
		IsInfected: false,
		Email:      email,
		Bio:        "",
	}
	_, err := db.ProfileCollection.Doc(id).Set(c, profile)
	if err != nil {
		logger.LogInternal(c, err, "error occurred when creating profile")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	registerResponse.Id = id
	registerResponse.Name = name
	registerResponse.Email = email

	c.IndentedJSON(http.StatusAccepted, registerResponse)

}