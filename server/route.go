package server

import (
	"expressd/controllers"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	orderGroup := router.Group("orders")
	{
		order := new(controllers.OrderController)
		orderGroup.GET("/:id", order.Retrieve)
	}
	return router
}
