package entity

//BlogEntity 数据库储存博客
type BlogEntity struct {
	Head     string `json:"head"`
	Number   uint   `gorm:"AUTO_INCREMENT;primary_key" json:"number"`
	Body     string `json:"body"`
	UserName string `json:"user_name"`
	Author   string `json:"author"`
	// Bytes    []byte `json:"bytes"`
}

//TableName 自定义表名称
func (BlogEntity) TableName() string {
	return "Blogs"
}
