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
		log.Print(err)
		panic(err)
	}

	// 自动生成表结构
	dbErr := Db.AutoMigrate(&entity.UserEntity{})
	if dbErr != nil {
		log.Print(dbErr)
		print(dbErr)
	}
	dbErr = Db.AutoMigrate(&entity.BlogEntity{})
	if dbErr != nil {
		log.Print(dbErr)
		print(dbErr)
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
func GetUserByLoginName(loginName string) (entity.UserEntity, bool) {
	var (
		user entity.UserEntity
	)
	err := Db.Where(&entity.UserEntity{LoginName: loginName}).First(&user).Error
	if err != nil {
		log.Print(err)
		return user, false
	}
	return user, true
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

//FindAllBlogs 获取blogs的所有数据
func FindAllBlogs() []*entity.BlogEntity {
	var blogs = make([]*entity.BlogEntity, 0)
	Db.Model(&entity.BlogEntity{}).Find(&blogs)
	return blogs
}

//FindAllUsers 获取users的所有数据
func FindAllUsers() []*entity.UserEntity {
	var users = make([]*entity.UserEntity, 0)
	Db.Model(&entity.UserEntity{}).Find(&users)
	return users
}

//GetBlogsByID 通过ID获取blogs的数据
func GetBlogsByID(number uint) (entity.BlogEntity, bool) {
	var blog = entity.BlogEntity{}
	err := Db.Where(&entity.BlogEntity{Number: number}).First(&blog).Error
	if err != nil {
		log.Print(err)
		return blog, false
	}
	return blog, true
}

// func Update(user userentity.UserEntity) {
// 	dbservice.Db.Model(&userentity.UserEntity{}).Where(&userentity.UserEntity{ID: user.ID}).Updates(user)
// }

//Delete 删除用户  TODO:有bug，会删除所有用户
func Delete(userID string) bool {
	// err := Db.Delete(&entity.UserEntity{}, userID).Error
	log.Print("待删除用户：" + userID)
	// err := Db.Model(&entity.UserEntity{}).Delete(&entity.UserEntity{LoginName: userID}).Error
	// if err != nil {
	// 	log.Print(err)
	// 	return false
	// }
	return true
}

//TODO: 修改用户信息 修改博客 删除博客
