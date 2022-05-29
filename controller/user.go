package controller

import (
	"blog-api/model"
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

func Create(c *gin.Context) {
	user := model.User{}
	err := c.Bind(&user)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad Request")
		return
	}
	userService := service.UserService{}
	err = userService.CreateUser(&user)
	if err != nil {
		c.String(http.StatusInternalServerError, "Server error")
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}
