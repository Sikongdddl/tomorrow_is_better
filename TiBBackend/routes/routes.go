package routes

import (
	"TiBBackend/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("")
	{
		api.POST("/register", controllers.RegisterUser)
		api.POST("/login", controllers.LoginUser)
		api.POST("/createTopic", controllers.CreateTopic)
		api.POST("/joinTopic", controllers.JoinTopic)
		api.GET("/getAllTopics", controllers.ListTopics)
	}
}
