package tools

import (
	"github.com/gin-contrib/sessions/cookie"
	// "github.com/gin-gonic/gin"
)

var sessionConfig map[string]interface{}

//SessionInit 初始化session,供main函数使用
func SessionInit() cookie.Store {
	sessionConfig := getSessionConfig()
	sessionStore := cookie.NewStore([]byte(sessionConfig["key"].(string)))
	return sessionStore
}

//getSessionConfig 获取依赖
func getSessionConfig() map[string]interface{} {
	sessionConfig = make(map[string]interface{})
	sessionConfig["key"] = "uselessblog"
	sessionConfig["loginname"] = "loginname"
	sessionConfig["nickname"] = "nickname"
	sessionConfig["age"] = 86400
	return sessionConfig
}

// func init() {
// sessionConfig = make(map[string]interface{})
// sessionConfig["key"] = "uselessblog"
// sessionConfig["name"] = "oj_session"
// sessionConfig["age"] = 86400
// sessionConfig["path"] = "/"
// }
