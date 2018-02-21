package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/konojunya/twblock-suspicious-account/service"
)

func IsAuthenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		if service.GetTwitterClient() == nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Next()
	}
}
