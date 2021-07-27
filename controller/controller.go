package controller

import (
	"log"
	"uselessBlog/model/usermodel"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	c.JSON(200, "hello")
}

func Login(c *gin.Context) {
	var user usermodel.User
	err := c.BindJSON(&user)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	c.JSON(200, "username: "+user.LoginName)
}

func Register(c *gin.Context) {
	var user usermodel.User
	err := c.BindJSON(&user)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	c.JSON(200, "username: "+user.LoginName)
}

func Del(c *gin.Context) {
	userID := c.Param(":id")
	c.JSON(200, "userID: "+userID)
}
