package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/konojunya/twblock-suspicious-account/service"
)

// GetUsers 怪しいアカウント一覧
func GetUsers(c *gin.Context) {
	users, err := service.GetUsers()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, users)
}

// BlockUser ユーザーをブロックする
func BlockUser(c *gin.Context) {
	id := c.Param("id")
	err := service.BlockUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.AbortWithStatus(http.StatusOK)
}
