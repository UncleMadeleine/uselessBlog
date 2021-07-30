package tools

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"uselessBlog/model/usermodel"

	"github.com/gookit/validate"
)

//EncodingSha256 一个hash函数
func EncodingSha256(data string) string {
	h := sha256.New()
	h.Write([]byte(data))
	sum := h.Sum(nil)
	s := hex.EncodeToString(sum)
	return string(s)
}

// UserForm struct
// type UserForm struct {
// 	Name     string    `validate:"required|minLen:7"`
// 	Email    string    `validate:"email" message:"email is invalid"`
// 	Age      int       `validate:"required|int|min:1|max:99" message:"int:age must int| min: age min value is 1"`
// 	CreateAt int       `validate:"min:1"`
// 	Safe     int       `validate:"-"`
// 	UpdateAt time.Time `validate:"required"`
// 	Code     string    `validate:"customValidator"` // 使用自定义验证器
// }
// UserForm struct
// type UserForm struct {
// 	LoginName string `validate:"required"`
// 	NickName  string `validate:"required"`
// 	Password  string `validate:"required|minLen:8"`
// 	Age       int    `validate:"required|int|min:1|max:99" message:"int:age must int| min: age min value is 1"`
// }

// // CustomValidator 定义在结构体中的自定义验证器
// func (f UserForm) CustomValidator(val string) bool {
// 	return len(val) == 4
// }

// // Messages  您可以自定义验证器错误消息
// func (f UserForm) Messages() map[string]string {
// 	return validate.MS{
// 		"required":      "oh! the {field} is required",
// 		"Name.required": "message for special field",
// 	}
// }

// // Translates 你可以自定义字段翻译
// func (f UserForm) Translates() map[string]string {
// 	return validate.MS{
// 		"Name":  "用户名称",
// 		"Email": "用户Email",
// 	}
// }

//Check 检验字段
func Check(u usermodel.User) bool {
	// u := &UserForm{
	// 	LoginName: "inhere",
	// }

	// 创建 Validation 实例
	v := validate.Struct(u)
	// v := validate.New(u)

	if v.Validate() { // 验证成功
		// do something ...
		return true
	}

	fmt.Println(v.Errors)               // 所有的错误消息
	fmt.Println(v.Errors.One())         // 返回随机一条错误消息
	fmt.Println(v.Errors.Field("Name")) // 返回该字段的错误消息
	return false
}
