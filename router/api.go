package router

import (
	"github.com/gin-gonic/gin"
	"github.com/konojunya/twblock-suspicious-account/controller"
)

func apiRouter(api *gin.RouterGroup) {
	api.GET("/users", controller.GetUsers)
	api.POST("/users/:screen_name/block", controller.BlockUser)
	api.GET("/healthcheck", controller.HealthCheck)
}
