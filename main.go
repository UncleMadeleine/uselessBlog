package main

import (
	"uselessBlog/dbservice"
	"uselessBlog/router"

	"github.com/gin-gonic/gin"
)

func main() {
	rout := gin.Default()
	dbservice.ConnectDb()
	router.Init(rout)
}
