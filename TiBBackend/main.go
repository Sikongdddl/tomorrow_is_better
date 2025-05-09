package main

import (
	"TiBBackend/config"
	"TiBBackend/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8081"},                   // 允许的源（前端的域名或端口）
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // 允许的请求方法
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // 允许的请求头
		AllowCredentials: true,                                                // 是否允许携带凭证（如 cookies）
	}))

	r.Static("/static/avatars", "./uploads/avatars")
	routes.SetupRoutes(r)

	r.Run(":8080") // listen and serve on localhost:8080
}
