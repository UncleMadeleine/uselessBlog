package tools

import (
	"github.com/gin-gonic/gin"
	// "google.golang.org/grpc/status"
)

//Spread 把后端错误传回前端
func Spread(c *gin.Context, state int, str1 string, str2 string) {
	errMessage := msg{
		Title:   str1,
		Content: str2,
	}
	c.JSON(state, errMessage)
}

//TypeReturn 返回正确的gin.H类型
func TypeReturn(data interface{}) gin.H {
	return gin.H{
		"status": 200,
		"msg":    "success!",
		"data":   data,
	}
}

type msg struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type returnType struct {
	Msg    string
	Data   interface{}
	Status int
}
