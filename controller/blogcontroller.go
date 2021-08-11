package controller

import (
	// "fmt"
	"log"
	"strconv"
	"time"
	"uselessBlog/dbservice"
	"uselessBlog/entity"

	// "uselessBlog/model"
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
	// fileName := "/localfiles/" + strconv.FormatInt(time.Now().Unix(), 10) + file.Filename
	fileName := "./localfiles/" + strconv.FormatInt(time.Now().Unix(), 10) + file.Filename
	// fileName := path.Join("/localfiles", strconv.FormatInt(time.Now().Unix(), 10)+file.Filename)
	err = c.SaveUploadedFile(file, fileName)
	if err != nil {
		log.Print(err)
		tools.Spread(c, 201, "储存文件失败", "储存文件失败")
		return
	}
	var ent entity.BlogEntity
	ent.Head = file.Filename
	ent.Head = ent.Head[:len(ent.Head)-3]
	// ent.Body = fileName[2:]
	ent.Body = fileName
	ent.UserName = userName.(string)
	blogName, ok := dbservice.UploadBlog(ent)
	if !ok {
		tools.Spread(c, 201, "数据库错误", "上传博客失败")
		return
	}
	tools.Spread(c, 200, "上传成功", "上传"+blogName+"博客成功")
}

//BlogLoad 动态加载博客
func BlogLoad(c *gin.Context, blogID string) {
	intID := tools.StrToUInt(blogID)
	if intID == 0 {
		tools.Spread(c, 201, "错误", "服务器错误,ID不为数字")
		return
	}
	blog, ok := dbservice.GetBlogsByID(intID)
	if !ok {
		tools.Spread(c, 201, "错误", "数据库错误")
		return
	}
	// log.Print("文件路径：" + blog.Body)
	data, ok := tools.ReadFile(blog.Body)
	if !ok {
		tools.Spread(c, 201, "错误", "读取文件错误")
	}
	c.Header("content-disposition", `attachment; filename=`+blog.Head)
	c.Data(200, "application/octet-stream", data)
}

//BlogIndex 获取首页所需的博客
func BlogIndex(c *gin.Context) {
	blogs := dbservice.FindAllBlogs()
	c.JSON(200, tools.TypeReturn(blogs))
}
