package main

import (
	"backend/app"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.New()
	app.MapRoutes(engine)
	engine.Run(":8080")
}
