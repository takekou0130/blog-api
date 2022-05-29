package main

import (
	"blog-api/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.GET("/users", controller.Index)
	engine.POST("/users", controller.Create)
	engine.Run(":3030")
}
