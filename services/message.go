package services

import (
	"cloud.google.com/go/firestore"
	"errors"
	"github.com/gin-gonic/gin"
	"konseki-be/db"
	"time"
)

const (
	INFECTED = "infected"
	CLOSE_CONTACT = "closeContact"
	DEFAULT = "default"
)


func CreateMessage(c *gin.Context, msgInfo MessageInfo) (*firestore.DocumentRef, *firestore.WriteResult, error) {
	// create messages in message collection
	// creation is done for everyone they have interacted with

	if isValidMsgType(msgInfo.MsgType) {
		err := errors.New("message type is not supported. please use a valid message type")
		return nil, nil, err
	}

	docRef, writeResult, err := db.MessageCollection.Add(c, msgInfo)
	return docRef, writeResult, err
}

func isValidMsgType(msgType string) bool {
	return msgType == INFECTED || msgType == CLOSE_CONTACT || msgType == DEFAULT
}

// isMessageImportant returns true if message is important
// importance is determined my message type as well as date of message creation
// if messages is "infected" importance last for 7 days after message created
// if message is "closeContact" importance last for 3 days
// importance of "infected" should outweigh "closeContact"
// if msgType is not "infected" or "closeContact", it is NOT important

func IsMessageImportant(msgType string, msgCreatedDate int64) bool {
	if msgType != "closeContact" && msgType != "infected" {
		return false
	}

	createdDate := time.Unix(msgCreatedDate, 0)
	elapseDateInfected := createdDate.AddDate(0,0,7)
	elapseDateCloseContact := createdDate.AddDate(0,0,3)
	currentDate := time.Now()

	if msgType == "infected" && currentDate.Before(elapseDateInfected){
		return true
	}

	if msgType == "closeContact" && currentDate.Before(elapseDateCloseContact){
		return true
	}

	return false
}


