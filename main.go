package main

import (
	"uselessBlog/dbservice"
	"uselessBlog/router"

	"github.com/gin-gonic/gin"
)

func main() {
	rout := gin.Default()
	rout.LoadHTMLGlob("static/html/*")
	// rout.Static("/static", "./static")
	dbservice.ConnectDb()
	router.Init(rout)
}
