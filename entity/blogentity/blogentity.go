package blogentity

type BlogEntity struct {
	Head   string
	Number string `gorm:"primaryKey"`
	Body   string
}

// 自定义表名称
func (BlogEntity) TableName() string {
	return "Blogs"
}
