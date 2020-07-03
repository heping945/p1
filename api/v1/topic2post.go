package api

import (
	"github.com/gin-gonic/gin"
	"pb/middleware"
	"pb/services"
)

func InitTopicRouter(Router *gin.RouterGroup) {
	TopicRouter := Router.Group("topic")
	{
		TopicRouter.GET("/", services.GetTopicList) // 获得列表
		TopicRouter.POST("/", middleware.Authen(), services.CreateTopic)
		TopicRouter.GET("/:slug", services.GetTopic)
		TopicRouter.DELETE("/:slug", middleware.Authen(), services.DeleteTopic)
	}
}

func InitPostRouter(Router *gin.RouterGroup) {
	PostRouter := Router.Group("post")
	{
		PostRouter.GET("/", services.GetPostList)                           // 获得列表
		PostRouter.POST("/", middleware.Authen(), services.CreatePost)      // 获得列表
		PostRouter.DELETE("/:id", middleware.Authen(), services.DeletePost) // 删除列表
	}
}

func InitTopicTagRouter(Router *gin.RouterGroup) {
	TagRouter := Router.Group("topictag")
	{
		TagRouter.GET("/", services.GetTopicTagList)

	}
}

func InitPostCategoryRouter(Router *gin.RouterGroup) {
	CategoryRouter := Router.Group("postcategory")
	{
		CategoryRouter.GET("/", services.GetPostCategoryList)
		CategoryRouter.POST("/", middleware.Authen(), services.CreatePostCategory)
		CategoryRouter.DELETE("/:name", middleware.Authen(), services.DeletePostCategory)
	}
}

func InitAuthenRouter(Router *gin.RouterGroup) {
	AuthenRouter := Router.Group("authen")
	{
		AuthenRouter.POST("/", services.CreateUser)
		AuthenRouter.GET("/:code", services.VerifyUser)
	}
}
