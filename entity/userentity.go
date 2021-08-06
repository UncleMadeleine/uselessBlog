package entity

//UserEntity 数据库储存用户
type UserEntity struct {
	Age       int
	Password  string
	LoginName string `gorm:"primaryKey"`
	NickName  string
	Admin     bool
}

//TableName 自定义表名称
func (UserEntity) TableName() string {
	return "Users"
}
