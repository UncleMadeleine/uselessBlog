package router

import (
	"uselessBlog/controller"

	"github.com/gin-gonic/gin"
)

// Init 初始化路由
func Init(router *gin.Engine) {
	home := router.Group("Home")
	router.RedirectFixedPath = true
	{
		home.GET("/", controller.Hello)
	}
	wqy := router.Group("user")
	{
		wqy.POST("/login", controller.Login)
		wqy.POST("/register", controller.Register)
		wqy.DELETE("/del/:id", controller.Del)
	}
	router.Run(":9000")

}
