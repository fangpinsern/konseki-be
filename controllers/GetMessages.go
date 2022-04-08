package controllers

import (
	"github.com/fangpinsern/konseki-be/db"
	"github.com/fangpinsern/konseki-be/logger"
	"github.com/fangpinsern/konseki-be/services"
	"github.com/fangpinsern/konseki-be/util"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/iterator"
	"net/http"
)

// GetMessageController only returns first 10 messages
func GetMessagesController(c *gin.Context) {
	userId := util.GetUserId(c)

	iter := db.MessageCollection.Where("Id", "==", userId).Documents(c)

	count := 0
	var messages []ResponseMessage
	for {
		doc, err := iter.Next()
		if err == iterator.Done || count > 10 {
			break
		}

		if err != nil {
			logger.LogInternal(c, err, "Message Data not found")
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		documentData := doc.Data()

		msgType := documentData["MsgType"].(string)
		msgCreatedDate := doc.CreateTime.Unix()

		msgIsImportant := services.IsMessageImportant(msgType, msgCreatedDate)

		msgInfo, err := db.UtilsCollection.Doc("messages").Get(c)

		msgInfoData := msgInfo.Data()
		msgVal, isMapContainsKey := msgInfoData[msgType]

		if !isMapContainsKey {
			msgVal = msgInfoData["default"]
		}

		newRes := ResponseMessage{
			Id:          documentData["Id"].(string),
			Message:     msgVal.(string),
			ExposureDate: documentData["ExposureDate"].(int64),
			MessageType: msgType,
			CreatedDate: msgCreatedDate,
			IsImportant: msgIsImportant,
		}

		messages = append(messages, newRes)
		count = count + 1

	}

	// headliner message - only have 1

	// sort and find headliner - infection > close contact > default
	var getMessagesResponse GetMessagesResponse
	getMessagesResponse.Messages = messages

	c.IndentedJSON(http.StatusAccepted, getMessagesResponse)


}
