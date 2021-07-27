package userentity

type userEntity struct {
	ID        string `gorm:"primaryKey"`
	Age       int
	Password  string
	LoginName string
	NickName  string
}

// 自定义表名称
func (userEntity) TableName() string {
	return "User"
}
