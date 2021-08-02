package controller

import (
	"log"
	"net/http"
	"uselessBlog/entity"
	"uselessBlog/model"
	"uselessBlog/service/userservice"
	"uselessBlog/tools"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//IndexTemplate : 加载主页
func IndexTemplate(c *gin.Context) {

	c.HTML(http.StatusOK, "index.html", gin.H{"msg": "successful!"})
	// c.JSON(200, "hello")
}

//LoginAPI : 登录操作
func LoginAPI(c *gin.Context) {
	var user model.User
	err := c.BindJSON(&user)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	if !tools.Check(user) {
		//返回前端报错
		return
	}
	ent := entity.UserEntity{}
	ent.LoginName = user.LoginName
	ent.Password = tools.EncodingSha256(user.Password)
	//TODO 验证登录
	en := userservice.GetUserByLoginName(ent.LoginName)
	if en.Password != ent.Password {
		//返回前端报错
		return
	}
	SaveSession(c, en)
	c.JSON(200, "username: "+user.LoginName)
}

//RegisterAPI : 注册操作
func RegisterAPI(c *gin.Context) {
	var user model.User
	err := c.BindJSON(&user)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	ent := entity.UserEntity{}
	ent.Age = user.Age
	ent.LoginName = user.LoginName
	ent.Password = tools.EncodingSha256(user.Password)
	userservice.SignIn(ent)
	c.JSON(200, "username: "+user.LoginName)
}

//UploadAPI 上传博客API
func UploadAPI(c *gin.Context) {
	UploadBlog(c)
}

//DelAPI : 删除数据库中的用户 TODO
func DelAPI(c *gin.Context) {
	userID := c.Param(":id")
	userservice.Delete(userID)
	c.JSON(200, "userName: "+userID)
}

//LoginTemplate : 加载登录页面
func LoginTemplate(c *gin.Context) {

	c.HTML(http.StatusOK, "login.html", gin.H{"msg": "successful!"})
	// c.JSON(200, "hello")
}

//UploadTemplate 加载上传界面
func UploadTemplate(c *gin.Context) {

	c.HTML(http.StatusOK, "uploadpage.html", gin.H{"msg": "successful!"})
	// c.JSON(200, "hello")
}

//RegisterTemplate : 加载注册页面
func RegisterTemplate(c *gin.Context) {

	c.HTML(http.StatusOK, "register.html", gin.H{"msg": "successful!"})
	// c.JSON(200, "hello")
}

//SaveSession : 登录时储存session
func SaveSession(c *gin.Context, us entity.UserEntity) {
	// 初始化session
	session := sessions.Default(c)
	// 保存session
	session.Set("Loginname", us.LoginName)
	session.Set("NickName", us.NickName)
	session.Set("Age", us.Age)
	session.Save()
	// 	// 取出session
	// 	if session.Get("Loginname") != nil {
	// 		//ops
	// 	}
}
