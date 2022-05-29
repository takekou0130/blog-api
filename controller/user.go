package controller

import (
	"blog-api/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	userService := service.UserService{}
	lists := userService.GetAllUsers()
	c.JSONP(http.StatusOK, gin.H{
		"message": "ok",
		"data":    lists,
	})
}
