package main

import (
	"backend/router"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.New()
	router.MapUrls(engine)
	engine.Run(":8080")
}
