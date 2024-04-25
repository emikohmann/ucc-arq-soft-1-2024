package controllers

import (
	"final/services"

	"github.com/gin-gonic/gin"
)

// Controller
func GetMessage(context *gin.Context) {
	message := context.Query("message")
	result := services.GetResult(message)
	context.JSON(200, result)
}
