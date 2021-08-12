package entity

//UserEntity 数据库储存用户
type UserEntity struct {
	Age       int `json:"age"`
	Password  string
	LoginName string `gorm:"primaryKey" json:"login_name"`
	NickName  string `json:"nick_name"`
	Admin     int
}

//TableName 自定义表名称
func (UserEntity) TableName() string {
	return "Users"
}
