package entity

//BlogEntity 数据库储存博客
type BlogEntity struct {
	Head     string
	Number   string `gorm:"primaryKey"`
	Body     string
	UserName string
}

//TableName 自定义表名称
func (BlogEntity) TableName() string {
	return "Blogs"
}
