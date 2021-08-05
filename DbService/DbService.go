package dbservice

import (
	"fmt"
	"log"
	"uselessBlog/entity"
	"uselessBlog/service/configservice"

	// mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"

	// "gorm.io/driver/mysql"
	"github.com/jinzhu/gorm"
)

//Db 数据库变量
var Db *gorm.DB

//ConnectDb 连接数据库
func ConnectDb() {
	var (
		err error
	)
	dbConfig := configservice.GetAppConfig("database")
	connectStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", dbConfig.USER,
		dbConfig.PASSWORD,
		dbConfig.HOST, dbConfig.NAME)
	Db, err = gorm.Open("mysql", connectStr)
	if err != nil {
		panic(err)
	}

	// 自动生成表结构
	dbErr := Db.AutoMigrate(&entity.UserEntity{})
	if dbErr != nil {
		println(err)
	}
	dbErr = Db.AutoMigrate(&entity.BlogEntity{})
	if dbErr != nil {
		print(err)
	}
}

//UploadBlog 数据库记录上传博客
func UploadBlog(ent entity.BlogEntity) (string, bool) {
	// log.Print(ent)
	err := Db.Create(&ent).Error
	if err != nil {
		log.Print(err)
		return "", false
	}
	// Db.Create(&user)
	return ent.Head, true
}

//GetUserByLoginName 以loginname获取用户
func GetUserByLoginName(loginName string) entity.UserEntity {
	var (
		user entity.UserEntity
	)
	Db.Where(&entity.UserEntity{LoginName: loginName}).First(&user)
	return user
}

//SignIn 注册新的用户
func SignIn(user entity.UserEntity) (string, bool) {
	// log.Print(user)
	log.Print(Db.NewRecord(user))
	err := Db.Create(&user).Error
	if err != nil {
		log.Print(err)
		return "", false
	}
	// Db.Create(&user)
	return user.LoginName, true
}

// func Update(user userentity.UserEntity) {
// 	dbservice.Db.Model(&userentity.UserEntity{}).Where(&userentity.UserEntity{ID: user.ID}).Updates(user)
// }

//Delete 删除用户  TODO
func Delete(userID string) {
	Db.Delete(&entity.UserEntity{}, userID)
}
