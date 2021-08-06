package tools

import "github.com/gin-gonic/gin"

//Spread 把后端错误传回前端
func Spread(c *gin.Context, str1 string, str2 string) {
	errMessage := msg{
		Title:   str1,
		Content: str2,
	}
	c.JSON(201, errMessage)
}

type msg struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}