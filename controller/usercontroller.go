package controller

import (
	"log"
	"uselessBlog/dbservice"
	"uselessBlog/entity"
	"uselessBlog/model"
	"uselessBlog/tools"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//LoginOperation 登录操作
func LoginOperation(c *gin.Context) {
	var user model.User
	err := c.BindJSON(&user)
	if err != nil {
		log.Print(err)
		tools.Spread(c, 201, "系统错误", "绑定json邦出问题了")
		return
	}
	ent := entity.UserEntity{}
	ent.LoginName = user.LoginName
	ent.Password = tools.EncodingSha256(user.Password)
	//TODO: 验证登录
	en, ok := dbservice.GetUserByLoginName(ent.LoginName)
	if !ok {
		tools.Spread(c, 201, "数据库错误", "可能根本没有这个用户")
		return
	}
	if en.Password != ent.Password {
		tools.Spread(c, 201, "登录失败", "用户名或密码错误")
		return
	}
	// LogOut(c)
	SaveSession(c, en)
	tools.Spread(c, 200, "登录成功", "登录成功")
}

//RegisterOperation 注册操作
func RegisterOperation(c *gin.Context) {
	var user model.User
	err := c.BindJSON(&user)
	if err != nil {
		log.Print(err)
		tools.Spread(c, 201, "系统错误", "绑定json邦出问题了")
		return
	}
	if !tools.Check(user) {
		tools.Spread(c, 201, "注册错误", "密码长度须不小于8")
		return
	}
	ent := entity.UserEntity{}
	ent.Age = tools.StrToInt(user.Age)
	if ent.Age == 0 {
		tools.Spread(c, 201, "年龄有误", "可能输入的不是数字")
	}
	ent.LoginName = user.LoginName
	ent.NickName = user.NickName
	ent.Admin = 1
	ent.Password = tools.EncodingSha256(user.Password)
	loginname, ok := dbservice.SignIn(ent)
	if !ok {
		tools.Spread(c, 201, "数据库错误", "可能是用户名重复了")
		return
	}
	// log.Print("注册操作正常，loginname：" + loginname)
	tools.Spread(c, 200, "注册成功", loginname+"注册成功")
	// c.JSON(200, "success")
}

//DelOperation 删除操作
func DelOperation(c *gin.Context) {
	// TODO 权限验证
	userID := c.Param(":id")
	ok := dbservice.Delete(userID)
	if !ok {
		tools.Spread(c, 201, "数据库错误", "删除失败")
		return
	}
	//TODO: 这里改成tools
	c.JSON(200, "userName: "+userID)
}

//SaveSession 登录时储存session
func SaveSession(c *gin.Context, us entity.UserEntity) {
	session := sessions.Default(c)
	session.Set("loginname", us.LoginName)
	session.Set("nickname", us.NickName)
	session.Set("age", us.Age)
	// LogOut(c)
	err := session.Save()
	if err != nil {
		log.Print("session储存失败")
		return
	}
	log.Print("储存成功，session储存用户：" + session.Get("loginname").(string))
	// // 取出session
	// if session.Get("Loginname") != nil {
	// 	//ops
	// }
}

//LogOut 注销登录
func LogOut(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	err := session.Save()
	if err != nil {
		log.Print(err)
		tools.Spread(c, 201, "出错了", "注销失败")
		return
	}
	// tools.Spread(c, 200, "注销成功", "注销成功")
}
