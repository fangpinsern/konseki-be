package services

import (
	"cloud.google.com/go/firestore"
	"github.com/fangpinsern/konseki-be/db"
	"github.com/fangpinsern/konseki-be/logger"
	"github.com/gin-gonic/gin"
)

func GetProfile(c *gin.Context, userId string) (Profile, error) {
	result, err := db.ProfileCollection.Doc(userId).Get(c)

	profileData := castProfileDocToProfile(result)
	if err != nil {
		return profileData, err
	}

	return profileData, nil
}

func CreateProfile(c *gin.Context, profileInfo Profile) (Profile, error){
	_, err := db.ProfileCollection.Doc(profileInfo.Id).Set(c, profileInfo)
	if err != nil {
		logger.LogInternal(c, err, "error occurred when creating profile")
		return Profile{}, err
	}
	return profileInfo, nil
}

func castProfileDocToProfile(profileDoc *firestore.DocumentSnapshot) Profile {
	documentData := profileDoc.Data()

	castedProfile := Profile{
		Id:         profileDoc.Ref.ID,
		Name:       documentData["Name"].(string),
		IsInfected: false,
		Email:      documentData["Email"].(string),
		Bio:        documentData["Bio"].(string),
	}

	return castedProfile
}