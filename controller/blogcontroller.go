package controller

import (
	"fmt"
	"strconv"
	"time"
	"uselessBlog/dbservice"
	"uselessBlog/entity"

	"github.com/gin-gonic/gin"
)

//UploadBlog 上传博客
func UploadBlog(c *gin.Context) {
	userName := c.PostForm("loginname")
	fmt.Println(userName)
	file, err := c.FormFile("file")
	if err != nil || userName == "" {
		//将错误返回前端 TODO
		return
	}
	//处理session TODO
	fileName := "/localfiles/" + strconv.FormatInt(time.Now().Unix(), 10) + file.Filename
	err = c.SaveUploadedFile(file, fileName)
	if err != nil {
		//将错误返回前端 TODO
		return
	}
	var ent entity.BlogEntity
	ent.Head = file.Filename
	ent.Body = fileName[1:]
	ent.Number = "1"
	ent.UserName = userName
	dbservice.UploadBlog(ent)
}
