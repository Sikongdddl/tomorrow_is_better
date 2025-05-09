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
		api.POST("/getTopicById", controllers.GetTopicById)

		api.POST("/getUserNameByID", controllers.GetNameByID)
		api.POST("/getUserInfoByID", controllers.GetUserInfoByID)
		api.POST("/uploadAvatar", controllers.UploadAvatar)
		api.POST("/leaveTopic", controllers.LeaveTopic)

		api.POST("/user/participated-topics", controllers.GetParticipatedTopics)
		api.POST("/user/created-topics", controllers.GetCreatedTopics)

		api.POST("/comment/list", controllers.GetCommentsByTopic)
		api.POST("/comment/add", controllers.AddComment)
		api.POST("/comment/delete", controllers.DeleteComment)

	}
}
