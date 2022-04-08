package services

import (
	"cloud.google.com/go/firestore"
	"github.com/fangpinsern/konseki-be/db"
	"github.com/fangpinsern/konseki-be/logger"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/iterator"
)

func GetAttendedEvents(c *gin.Context, userId string) ([]Event, error) {
	var responseEventList []Event
	iter := db.EventCollection.Where("Attended", "array-contains", userId).Documents(c)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			logger.LogInternal(c, err, "Event Data not found")
			return nil, err
		}

		newEvent := castEventDocToEvent(doc)

		responseEventList = append(responseEventList, newEvent)
	}

	return responseEventList, nil
}

func castEventDocToEvent(eventDoc *firestore.DocumentSnapshot) Event {
	documentData := eventDoc.Data()

	var attendedIds []string

	for _, val := range documentData["Attended"].([]interface{}) {
		attendedIds = append(attendedIds, val.(string))
	}
	newEvent := Event{
		Id:        eventDoc.Ref.ID,
		Name:      documentData["Name"].(string),
		Attended:  attendedIds,
		Date:      documentData["Date"].(int64),
		CreatorId: documentData["CreatorId"].(string),
	}

	return newEvent
}

func JoinEvent(c *gin.Context, eventId, userId string) (Event, error) {
	result, err := db.EventCollection.Doc(eventId).Get(c)
	eventData := castEventDocToEvent(result)
	if err != nil {
		//joinEventResponse.IsSuccess = false
		//c.JSON(http.StatusBadRequest, joinEventResponse)
		return eventData, err
	}

	// check if the user is already in the event
	attendedIds := append(eventData.Attended, userId)

	_, err1 := db.EventCollection.Doc(eventId).Update(c, []firestore.Update{{Path: "Attended", Value: attendedIds}})
	if err1 != nil {
		logger.LogInternal(c, err1, "error when updated event with attended data")
		return eventData, err1
	}

	updatedEventDoc, err := db.EventCollection.Doc(eventId).Get(c)
	returnCastEvent := castEventDocToEvent(updatedEventDoc)


	return returnCastEvent, nil

}
