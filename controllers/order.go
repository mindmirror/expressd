package controllers

import (
	"expressd/models"

	"github.com/gin-gonic/gin"
)

type OrderController struct{}

var orderModel = new(models.Order)

func (o OrderController) Retrieve(c *gin.Context) {
	c.JSON(200, gin.H{"message": "ok"})
}
