package userservice

import (
	"uselessBlog/dbservice"
	"uselessBlog/entity/userentity"
)

func ByLoginNameGetUser(loginName string) userentity.UserEntity {
	var (
		user userentity.UserEntity
	)
	dbservice.Db.Where(&userentity.UserEntity{LoginName: loginName}).First(&user)
	return user
}

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

func Delete(userId string) {
	dbservice.Db.Delete(&userentity.UserEntity{}, userId)
}
