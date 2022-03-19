package util

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"strings"
)

const (
	CorrelationIDContextKey = "correlation_id"
	AuthContextKey          = "user"
	UserIdContextKey = "userId"
	UserEmailContextKey = "userEmail"
	UserAuthTimeContextKey = "userAuthTime"
)

func GetCorrelationID(c *gin.Context) string {
	value, exist := c.Get(CorrelationIDContextKey)

	if !exist {
		return ""
	}

	valueString := value.(string)

	return valueString
}

func SetCorrelationID(c *gin.Context) {
	uuidWithHyphen := uuid.New()
	token := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	c.Set(CorrelationIDContextKey, token)
}

func GetPath(c *gin.Context) string {
	value := c.Request.RequestURI

	return value
}

func GetMethod(c *gin.Context) string {
	value := c.Request.Method
	return value
}

func SetUser(c *gin.Context, user map[string]interface{}) {
	c.Set(AuthContextKey, user)
	c.Set(UserIdContextKey, user["user_id"])
	c.Set(UserEmailContextKey,user["email"])
	c.Set(UserAuthTimeContextKey, user["auth_time"])
}

func GetUserId (c *gin.Context) string {
	value, exist := c.Get(UserIdContextKey)
	if !exist {
		return ""
	}

	valueString := value.(string)

	return valueString
}

func GetUserEmail (c *gin.Context) string {
	value, exist := c.Get(UserEmailContextKey)
	if !exist {
		return ""
	}

	valueString := value.(string)

	return valueString
}

func GetUserAuthTime (c *gin.Context) int {
	value, exist := c.Get(UserAuthTimeContextKey)
	if !exist {
		return 0
	}

	valueString := value.(int)

	return valueString
}
