// package main

// import "github.com/gin-gonic/gin"

// func main() {
// 	r := gin.Default()
// 	r.GET("/hi", func(c *gin.Context) {
// 		c.JSON(200, "hello world")
// 	})
// 	r.Run(":8888") // 监听并在 127.0.0.1:8888 上启动服务
// }
package main

import (
	"uselessBlog/Routers"
	"uselessBlog/Service/DbService"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	DbService.ConnectDb()
	Routers.Init(router)
}
