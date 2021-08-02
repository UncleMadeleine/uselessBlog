package entity

type UserEntity struct {
	Age       int
	Password  string
	LoginName string `gorm:"primaryKey"`
	NickName  string
}

// 自定义表名称
func (UserEntity) TableName() string {
	return "Users"
}
