package main

import (
	"TiBBackend/config"
	"TiBBackend/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()
	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run(":8080") // listen and serve on localhost:8080
}
