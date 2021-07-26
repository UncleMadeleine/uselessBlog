package LoginControoler

import (
	"uselessBlog/Model/ViewModel/ApiState"
	"uselessBlog/Service/UserService"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	loginName := c.PostForm("loginName")
	password := c.PostForm("password")
	if loginName == "" || password == "" {
		ApiState.ArgErrApiResult(c, "loginName or password")
		return
	}
	userInfo := UserService.GetUserInfo(loginName, password)
	if userInfo.Id == "" {
		ApiState.ArgErrApiResult(c, "loginName or password")
		return
	}
	//cookieId := UtilService.UUId()
	//CacheService.SetCache(cookieId, userInfo)

	ApiState.ResponseSuccess(c, userInfo)
}
