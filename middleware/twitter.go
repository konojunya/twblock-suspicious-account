package middleware

import (
	"net/http"

	"github.com/konojunya/twblock-suspicious-account/service/twitter"

	"github.com/gin-gonic/gin"
)

func IsAuthenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		if twitter.GetClient() == nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Next()
	}
}
