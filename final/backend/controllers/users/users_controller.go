package users

import (
	usersDomain "backend/domain/users"
	usersService "backend/services/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(context *gin.Context) {
	var loginRequest usersDomain.LoginRequest
	context.BindJSON(&loginRequest)
	response := usersService.Login(loginRequest)
	context.JSON(http.StatusOK, response)
}
