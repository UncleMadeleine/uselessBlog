package userservice

import (
	"uselessBlog/dbservice"
	"uselessBlog/entity/userentity"
)

//ByLoginNameGetUser 以loginname获取用户
func ByLoginNameGetUser(loginName string) userentity.UserEntity {
	var (
		user userentity.UserEntity
	)
	dbservice.Db.Where(&userentity.UserEntity{LoginName: loginName}).First(&user)
	return user
}

//SignIn 注册新的用户
func SignIn(user userentity.UserEntity) string {
	if dbservice.Db.NewRecord(user) {
		dbservice.Db.Create(&user)
	}
	// dbservice.Db.Create(&user)
	return user.LoginName
}

// func Update(user userentity.UserEntity) {
// 	dbservice.Db.Model(&userentity.UserEntity{}).Where(&userentity.UserEntity{ID: user.ID}).Updates(user)
// }

//Delete 删除用户  TODO
func Delete(userID string) {
	dbservice.Db.Delete(&userentity.UserEntity{}, userID)
}
