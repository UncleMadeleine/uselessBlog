package controller

import (
	"log"
	"net/http"
	"uselessBlog/entity/userentity"
	"uselessBlog/model/usermodel"
	"uselessBlog/service/userservice"
	"uselessBlog/tools"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {

	c.HTML(http.StatusOK, "index.html", gin.H{"msg": "successful!"})
	// c.JSON(200, "hello")
}

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
	userservice.ByLoginNameGetUser(ent.LoginName)
	c.JSON(200, "username: "+user.LoginName)
}

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

func Del(c *gin.Context) {
	userID := c.Param(":id")
	userservice.Delete(userID)
	c.JSON(200, "userName: "+userID)
}
