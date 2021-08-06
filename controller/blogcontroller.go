package controller

import (
	"fmt"
	"log"
	"strconv"
	"time"
	"uselessBlog/dbservice"
	"uselessBlog/entity"
	"uselessBlog/tools"

	"github.com/gin-gonic/gin"
)

//UploadBlog 上传博客
func UploadBlog(c *gin.Context) {
	// log.Print("正在上传")
	userName := c.PostForm("loginname")
	fmt.Println(userName)
	file, err := c.FormFile("file")
	if err != nil || userName == "" {
		log.Print(err)
		tools.Spread(c, "出错了", "文件有误或未登录")
		return
	}
	//处理session TODO
	fileName := "/localfiles/" + strconv.FormatInt(time.Now().Unix(), 10) + file.Filename
	err = c.SaveUploadedFile(file, fileName)
	if err != nil {
		log.Print(err)
		tools.Spread(c, "储存文件失败", "储存文件失败")
		return
	}
	var ent entity.BlogEntity
	ent.Head = file.Filename
	ent.Body = fileName[1:]
	ent.UserName = userName
	dbservice.UploadBlog(ent)
}
