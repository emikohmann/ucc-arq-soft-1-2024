package main

import (
	"final/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	// ej. http://localhost:3000/test?message=hola
	engine := gin.New()
	engine.GET("/test", controllers.GetMessage)
	engine.Run(":3000")
}
