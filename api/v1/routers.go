package api

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"log"
	_ "pb/docs"
	"pb/middleware"
)

func Routers() *gin.Engine {
	var Router = gin.Default()
	Router.Use(middleware.Cors())

	log.Println("use middleware cors")
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	log.Println("register swagger handler")
	ApiGroup := Router.Group("api/v1")
	{
		InitTopicTagRouter(ApiGroup)
		InitPostRouter(ApiGroup)
		InitTopicRouter(ApiGroup)
		InitPostCategoryRouter(ApiGroup)
		InitAuthenRouter(ApiGroup)
	}

	log.Println("router register success")
	return Router
}
