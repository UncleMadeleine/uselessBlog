package usermodel

//用户表单
type User struct {
	ID        string `form:"id"`
	Age       int    `form:"age"`
	Password  string `form:"password"`
	LoginName string `form:"loginname"`
	NickName  string `form:"nickname"`
}
