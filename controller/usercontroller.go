package controller

import (
	"log"
	"strconv"
	"uselessBlog/entity"
	"uselessBlog/model"

	// "uselessBlog/service/userservice"
	"uselessBlog/dbservice"
	"uselessBlog/tools"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//LoginOperation 登录操作
func LoginOperation(c *gin.Context) {
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
	en := dbservice.GetUserByLoginName(ent.LoginName)
	if en.Password != ent.Password {
		//返回前端报错
		return
	}
	SaveSession(c, en)
	c.JSON(200, "username: "+user.LoginName)
}

//RegisterOperation 注册操作
func RegisterOperation(c *gin.Context) {
	var user model.User
	err := c.BindJSON(&user)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	ent := entity.UserEntity{}
	intUserAge, err := strconv.Atoi(user.Age)
	if err != nil {
		//返回错误给前端
		c.JSON(201, "年龄不是数字")
		return
	}
	ent.Age = intUserAge
	ent.LoginName = user.LoginName
	ent.NickName = user.NickName
	ent.Password = tools.EncodingSha256(user.Password)
	loginname, ok := dbservice.SignIn(ent)
	if !ok {
		//返回错误给前端
		tools.Spread(c, "数据库错误", "可能是用户名重复了")
		return
	}
	log.Print("注册操作正常，loginname：" + loginname)
	c.JSON(200, "username: "+loginname)
	// c.JSON(200, "success")
}

//DelOperation 删除操作
func DelOperation(c *gin.Context) {
	// TODO
	userID := c.Param(":id")
	dbservice.Delete(userID)
	c.JSON(200, "userName: "+userID)
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
