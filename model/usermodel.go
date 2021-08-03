package model

//User 用户表单
type User struct {
	// ID        string `form:"id"`
	NickName  string `form:"nickname" validate:"required"`
	Password  string `form:"password" validate:"required"`
	LoginName string `form:"loginname" validate:"required|minLen:8"`
	Age       string `form:"age" validate:"required"`
	// Age  int `form:"age" validate:"required|int|min:1|max:99" message:"int:age must int| min: age min value is 1"`
}
