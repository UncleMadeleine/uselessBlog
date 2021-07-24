package Entity

import "gorm.io/gorm"

type UserEntity struct {
	gorm.Model
	Age       int    `form:"age"`
	Password  string `form:"password"`
	LoginName string `form:"loginName"`
	NickName  string `form:"nickName"`
}

func (UserEntity) TableName() string {
	return "User"
}
