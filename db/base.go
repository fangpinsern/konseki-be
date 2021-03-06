package db

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/fangpinsern/konseki-be/config"
	"google.golang.org/api/option"
)

var Database *firebase.App

var AuthClient *auth.Client

var DatabaseClient *firestore.Client

var ProfileCollection *firestore.CollectionRef
var EventCollection *firestore.CollectionRef
var MessageCollection *firestore.CollectionRef
var InfectionCollection *firestore.CollectionRef
var UtilsCollection *firestore.CollectionRef


func InitializeDatabase(c context.Context) {

	opt:=option.WithCredentialsJSON(config.GetFirestoreCreds())

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic("Failed to initialize DB")
	}

	Database = app

	DatabaseClient, err = app.Firestore(c)
	if err != nil {
		panic(err)
	}

	ProfileCollection = DatabaseClient.Collection("profile")
	EventCollection = DatabaseClient.Collection("event")
	MessageCollection = DatabaseClient.Collection("message")
	InfectionCollection = DatabaseClient.Collection("infection")
	UtilsCollection = DatabaseClient.Collection("utils")

	AuthClient ,_ = app.Auth(c)

}


