package dbservice

import (
	"fmt"
	"uselessBlog/entity"
	"uselessBlog/service/configservice"

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
func UploadBlog(ent entity.BlogEntity) string {
	if Db.NewRecord(ent) {
		Db.Create(&ent)
	}
	// dbservice.Db.Create(&user)
	return ent.Head
}
