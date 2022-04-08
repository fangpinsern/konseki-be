package controllers

import (
	"github.com/fangpinsern/konseki-be/db"
	"github.com/fangpinsern/konseki-be/logger"
	"github.com/fangpinsern/konseki-be/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetProfileInfoController(c *gin.Context) {
	userId := util.GetUserId(c)

	var getProfileInfoResponse GetProfileInfoResponse

	result, err := db.ProfileCollection.Doc(userId).Get(c)
	if err != nil {
		logger.LogInternal(c, err, "profile data not found")
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	profileData := result.Data()

	getProfileInfoResponse.Id = result.Ref.ID
	getProfileInfoResponse.Name = profileData["Name"].(string)
	getProfileInfoResponse.Email = profileData["Email"].(string)
	getProfileInfoResponse.Bio = profileData["Bio"].(string)

	c.IndentedJSON(http.StatusAccepted, getProfileInfoResponse)

}
