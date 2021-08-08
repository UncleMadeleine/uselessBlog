package controller

import (
	"log"
	"strconv"
	"time"
	"uselessBlog/dbservice"
	"uselessBlog/entity"
	"uselessBlog/tools"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//UploadBlog 上传博客
func UploadBlog(c *gin.Context) {
	// log.Print("正在上传")
	session := sessions.Default(c)
	if session.Get("loginname") == nil {
		tools.Spread(c, 201, "出错了", "未登录")
	}
	userName := session.Get("loginname")
	// log.Print("session中储存的用户:" + userName.(string))
	file, err := c.FormFile("file")
	if err != nil {
		log.Print(err)
		tools.Spread(c, 201, "出错了", "文件有误")
		return
	}
	fileName := "/localfiles/" + strconv.FormatInt(time.Now().Unix(), 10) + file.Filename
	// err = c.SaveUploadedFile(file, fileName)
	// if err != nil {
	// 	log.Print(err)
	// 	tools.Spread(c, 201, "储存文件失败", "储存文件失败")
	// 	return
	// }
	var ent entity.BlogEntity
	ent.Head = file.Filename
	ent.Body = fileName[1:]
	ent.UserName = userName.(string)
	blogName, ok := dbservice.UploadBlog(ent)
	if !ok {
		tools.Spread(c, 201, "数据库错误", "上传博客失败")
		return
	}
	tools.Spread(c, 200, "上传成功", "上传"+blogName+"博客成功")
}
