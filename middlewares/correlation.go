package middlewares

import (
	"github.com/gin-gonic/gin"
	"konseki-be/util"
)

func CorrelationData() gin.HandlerFunc {
	return func(c *gin.Context) {
		util.SetCorrelationID(c)
	}
}
