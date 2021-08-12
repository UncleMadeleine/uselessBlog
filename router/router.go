package router

import (
	"uselessBlog/controller"

	"github.com/gin-gonic/gin"
)

//Init 初始化路由
func Init(router *gin.Engine) {
	home := router.Group("/")
	router.RedirectFixedPath = true
	{
		home.GET("/", controller.IndexTemplate)
		home.POST("/", controller.IndexAPI)
		home.GET("/login", controller.LoginTemplate)
		home.GET("/register", controller.RegisterTemplate)
		home.GET("/upload", controller.UploadTemplate)
		home.GET("/admin", controller.AdminTemplate)
		home.POST("/admin", controller.AdminAPI)
	}
	wqy := router.Group("user")
	{
		wqy.POST("/login", controller.LoginAPI)
		wqy.POST("/register", controller.RegisterAPI)
		wqy.GET("/del/:id", controller.DelAPI)
		// wqy.DELETE("/del/:id", controller.DelAPI)
		wqy.POST("/upload", controller.UploadAPI)
	}
	ljw := router.Group("blog")
	{
		ljw.GET("/:id", controller.BlogMedalTemplate)
		ljw.POST("/:id", controller.BlogMedalAPI)
	}
	router.Run(":9000")

}
