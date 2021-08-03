package main

import (
	"uselessBlog/dbservice"
	"uselessBlog/router"
	"uselessBlog/tools"

	"github.com/gin-gonic/gin"
)

func main() {
	tools.LogInit()
	rout := gin.Default()
	rout.LoadHTMLGlob("static/html/*")
	rout.Static("/static", "./static")
	dbservice.ConnectDb()
	router.Init(rout)
}
