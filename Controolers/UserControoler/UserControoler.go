package UserControoler

import (
	"uselessBlog/Model/ViewModel/ApiState"
	"uselessBlog/Service/UserService"

	"github.com/gin-gonic/gin"
)

func GetUserInfo(c *gin.Context) {
	loginName := c.Query("loginName")
	password := c.Query("password")
	if loginName == "" || password == "" {
		ApiState.ArgErrApiResult(c, "loginName or password")
		return
	}
	userInfo := UserService.GetUserInfo(loginName, password)
	ApiState.ResponseSuccess(c, userInfo)
}
