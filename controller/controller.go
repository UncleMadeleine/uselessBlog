package controller

import (
	"log"
	"net/http"
	"uselessBlog/entity/userentity"
	"uselessBlog/model/usermodel"
	"uselessBlog/service/userservice"
	"uselessBlog/tools"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//Hello : 加载主页
func Hello(c *gin.Context) {

	c.HTML(http.StatusOK, "index.html", gin.H{"msg": "successful!"})
	// c.JSON(200, "hello")
}

//Login : 登录操作
func Login(c *gin.Context) {
	var user usermodel.User
	err := c.BindJSON(&user)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	ent := userentity.UserEntity{}
	ent.Age = user.Age
	ent.LoginName = user.LoginName
	ent.Password = tools.EncodingSha256(user.Password)
	//TODO 验证登录
	en := userservice.ByLoginNameGetUser(ent.LoginName)
	SaveSession(c, en)
	c.JSON(200, "username: "+user.LoginName)
}

//Register : 注册操作
func Register(c *gin.Context) {
	var user usermodel.User
	err := c.BindJSON(&user)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	ent := userentity.UserEntity{}
	ent.Age = user.Age
	ent.LoginName = user.LoginName
	ent.Password = tools.EncodingSha256(user.Password)
	userservice.SignIn(ent)
	c.JSON(200, "username: "+user.LoginName)
}

//Del : 删除数据库中的用户 TODO
func Del(c *gin.Context) {
	userID := c.Param(":id")
	userservice.Delete(userID)
	c.JSON(200, "userName: "+userID)
}

//HelloLogin : 加载登录页面
func HelloLogin(c *gin.Context) {

	c.HTML(http.StatusOK, "login.html", gin.H{"msg": "successful!"})
	// c.JSON(200, "hello")
}

//HelloRegister : 加载注册页面
func HelloRegister(c *gin.Context) {

	c.HTML(http.StatusOK, "register.html", gin.H{"msg": "successful!"})
	// c.JSON(200, "hello")
}

//SaveSession : 登录时储存session
func SaveSession(c *gin.Context, us userentity.UserEntity) {
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
