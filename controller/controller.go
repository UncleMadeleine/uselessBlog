package controller

import (
	"net/http"

	// "uselessBlog/entity"

	// "github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//IndexTemplate : 加载主页
func IndexTemplate(c *gin.Context) {
	// blogs := dbservice.FindAllBlogs()

	// c.HTML(http.StatusOK, "index.html", tools.TypeReturn(blogs))
	c.HTML(http.StatusOK, "index.html", gin.H{"msg": "success!"})
	// c.JSON(200, "hello")
}

//IndexAPI 获取首页所需的博客
func IndexAPI(c *gin.Context) {
	BlogIndex(c)
}

//BlogMedalTemplate 加载各博客页面
func BlogMedalTemplate(c *gin.Context) {
	blogID := c.Param(":id")
	c.JSON(200, blogID)
	c.HTML(http.StatusOK, "Medal.html", gin.H{"msg": "success!"})
	// c.JSON(200, "hello")
}

//BlogMedalAPI 便于js动态加载博客内容
func BlogMedalAPI(c *gin.Context) {
	blogID := c.Param(":id")
	BlogLoad(c, blogID)
	// c.JSON(200, "hello")
}

//LoginAPI : 登录操作
func LoginAPI(c *gin.Context) {
	LoginOperation(c)
}

//RegisterAPI : 注册操作
func RegisterAPI(c *gin.Context) {
	RegisterOperation(c)
}

//UploadAPI 上传博客API
func UploadAPI(c *gin.Context) {
	UploadBlog(c)
}

//DelAPI : 删除数据库中的用户 TODO
func DelAPI(c *gin.Context) {
	DelOperation(c)
}

//LoginTemplate : 加载登录页面
func LoginTemplate(c *gin.Context) {

	c.HTML(http.StatusOK, "login.html", gin.H{"msg": "successful!"})
}

//UploadTemplate 加载上传界面
func UploadTemplate(c *gin.Context) {
	c.HTML(http.StatusOK, "uploadpage.html", gin.H{"msg": "successful!"})
}

//RegisterTemplate : 加载注册页面
func RegisterTemplate(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{"msg": "successful!"})
}
