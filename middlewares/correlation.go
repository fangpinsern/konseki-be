package middlewares

import (
	"github.com/fangpinsern/konseki-be/util"
	"github.com/gin-gonic/gin"
)

func CorrelationData() gin.HandlerFunc {
	return func(c *gin.Context) {
		util.SetCorrelationID(c)
	}
}
