package dbservice

import (
	"fmt"
	"uselessBlog/entity/userentity"
	"uselessBlog/service/configservice"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	// "gorm.io/driver/mysql"
	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

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
	dbErr := Db.AutoMigrate(&userentity.UserEntity{})
	if dbErr != nil {
		println(err)
	}
}
