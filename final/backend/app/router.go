package app

import (
	"backend/controllers/users"

	"github.com/gin-gonic/gin"
)

func MapRoutes(engine *gin.Engine) {
	engine.POST("/users/login", users.Login)
}
