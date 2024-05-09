package main

import (
	"backend/router"

	"backend/clients"

	"github.com/gin-gonic/gin"
)

func main() {
	clients.StartDB()

	engine := gin.New()
	router.MapUrls(engine)
	engine.Run(":8080")
}
