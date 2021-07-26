package RegisterService

import (
	"uselessBlog/Model/Entity"
	"uselessBlog/Service/DbService"
)

// 注册用户
func RegisterUser(user Entity.UserEntity) string {
	user.Id = "123"
	DbService.Db.Create(&user)
	return user.Id
}
