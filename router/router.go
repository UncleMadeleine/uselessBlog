package router

import (
	"uselessBlog/controller"

	"github.com/gin-gonic/gin"
)

// Init 初始化路由
func Init(router *gin.Engine) {
	home := router.Group("/")
	router.RedirectFixedPath = true
	{
		home.GET("/", controller.IndexTemplate)
		home.GET("/login", controller.LoginTemplate)
		home.GET("/register", controller.RegisterTemplate)
		home.GET("/upload", controller.UploadTemplate)

	}
	wqy := router.Group("user")
	{
		wqy.POST("/login", controller.LoginAPI)
		wqy.POST("/register", controller.RegisterAPI)
		wqy.DELETE("/del/:id", controller.DelAPI)
		home.POST("/upload", controller.UploadAPI)
	}
	router.Run(":9000")

}
