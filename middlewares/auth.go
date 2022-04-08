package middlewares

import (
	"fmt"
	"github.com/fangpinsern/konseki-be/db"
	"github.com/fangpinsern/konseki-be/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthorizeToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer"
		authHeader := c.GetHeader("Authorization")

		if len(authHeader) <= 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Username or Password"})
			return
		}

		tokenString := authHeader[len(BEARER_SCHEMA)+1:]
		token, err := db.AuthClient.VerifyIDToken(c, tokenString)
		if err != nil {
			fmt.Println(err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Username or Password"})
			return
		}

		util.SetUser(c, token.Claims)

	}
}

