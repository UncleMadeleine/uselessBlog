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
		home.GET("/login", controller.LoginTemplate)
		home.GET("/register", controller.RegisterTemplate)
		home.GET("/upload", controller.UploadTemplate)

	}
	wqy := router.Group("user")
	{
		wqy.POST("/login", controller.LoginAPI)
		wqy.POST("/register", controller.RegisterAPI)
		wqy.DELETE("/del/:id", controller.DelAPI)
		wqy.POST("/upload", controller.UploadAPI)
	}
	ljw := router.Group("blog")
	{
		//TODO:
		// dbvalues := dbservice.FindBlogs()
		ljw.GET("/:id", controller.BlogMedalTemplate)
		ljw.GET("/API/:id", controller.BlogMedalAPI)
	}
	router.Run(":9000")

}
