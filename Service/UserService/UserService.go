package userservice

import (
	"uselessBlog/dbservice"
	"uselessBlog/model/usermodel"
)

func ByLoginNameGetUser(loginName string) usermodel.User {
	var (
		user usermodel.User
	)
	dbservice.Db.Where(&usermodel.User{LoginName: loginName}).First(&user)
	return user
}

func SignIn(user usermodel.User) string {
	user.ID = "123"
	dbservice.Db.Create(&user)
	return user.ID
}

func Update(user usermodel.User) {
	dbservice.Db.Model(&usermodel.User{}).Where(&usermodel.User{ID: user.ID}).Updates(user)
}

func Delete(userId string) {
	dbservice.Db.Delete(&usermodel.User{}, userId)
}
