package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"konseki-be/config"
	"konseki-be/controllers"
	"konseki-be/db"
	"konseki-be/logger"
	"konseki-be/middlewares"
	"konseki-be/util"
	"net/http"
	"os"
)

type PingStuct struct {
	Message string `json:"message"`
	UserId string `json:"user_id"`
	Email string `json:"email"`
}

func main() {
	config.InitializeConfig("./env")
	// may need to initialize db
	db.InitializeDatabase(context.Background())
	// need to initialize email service
	// initialize logger
	logger.InitializeLogger()
	router := gin.Default()

	router.Use(middlewares.CorrelationData())
	router.Use(middlewares.AuthorizeToken())

	router.GET("/ping", healthCheck)

	router.POST("/register", controllers.RegisterController)
	router.GET("/profile", controllers.GetProfileInfoController)
	router.POST("/event/create", controllers.CreateEventController)
	router.GET("/event/all", controllers.GetEventsController)
	router.POST("/event/join", controllers.JoinEventController)

	router.POST("/infection/update", controllers.UpdateStatusController)

	router.GET("/messages/all", controllers.GetMessagesController)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	router.Run("0.0.0.0:" + PORT)
}

func healthCheck(c *gin.Context) {
	msg := PingStuct{
		Message: "helloworld",
		UserId: util.GetUserId(c),
		Email: util.GetUserEmail(c),
	}
	c.IndentedJSON(http.StatusOK, msg)
}
