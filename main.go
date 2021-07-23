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
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H {
			"Name": "Hello world !",
		})
	})
	router.Run(": 8888")
}