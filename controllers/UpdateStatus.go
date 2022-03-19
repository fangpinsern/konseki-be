package controllers

import (
	"github.com/gin-gonic/gin"
	"konseki-be/db"
	"konseki-be/logger"
	"konseki-be/services"
	"konseki-be/util"
	"net/http"
)

func UpdateStatusController(c *gin.Context) {
	var updateStatusRequest UpdateStatusRequest
	var updateStatusResponse UpdateStatusResponse

	if err := c.BindJSON(&updateStatusRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// create a new collection to store all infections
	userId := util.GetUserId(c)
	infectionDate := updateStatusRequest.Date

	infection := Infections{
		UserId: userId,
		Date:   infectionDate,
	}

	result, _, err := db.InfectionCollection.Add(c, infection)
	if err != nil {
		logger.LogInternal(c, err, "Something went wrong when adding infection")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}

	infectionMessage := services.MessageInfo{
		Id:      userId,
		ExposureDate:    infectionDate,
		MsgType: "infected",
	}

	_, _, err = services.CreateMessage(c, infectionMessage)
	if err != nil {
		logger.LogInternal(c, err, "Error occurred when adding infection message")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}

	updateStatusResponse.Id = result.ID
	updateStatusResponse.IsSuccess = true
	updateStatusResponse.Date = infectionDate

	// TODO: create messages for everyone - message changes depending on days of interaction
	// future: create separate service to use go routines
	// get everyone that has interacted with this person
	// create message for everyone
	// notify the people

	//iter := db.EventCollection.Where("Attended", "array-contains", userId).Documents(c)
	attendedEvents, err := services.GetAttendedEvents(c, userId)

	var attendedInfo []services.MessageInfo
	alreadyAdded := map[string]bool{}
	for _, doc := range attendedEvents {
		for _, val := range doc.Attended {
			castedVal := val
			if castedVal == userId {
				continue
			}
			_, isMapContainsKey := alreadyAdded[castedVal]
			if !isMapContainsKey {
				messageInfo := services.MessageInfo{
					Id:   castedVal,
					ExposureDate: doc.Date,
					MsgType: "closeContact",
				}
				attendedInfo = append(attendedInfo, messageInfo)
				alreadyAdded[castedVal] = true
			}
		}
	}

	for _, messageVal := range attendedInfo {
		_, _, err := services.CreateMessage(c, messageVal)
		if err != nil {
			logger.LogInternal(c, err, "Error occurred when creating message for" + messageVal.Id)
		}
	}

	c.IndentedJSON(http.StatusAccepted, updateStatusResponse)
	return
}