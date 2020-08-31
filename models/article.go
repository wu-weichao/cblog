package models

type Article struct {
	Model

	Tags        string `gorm:"comment:关联标签" json:"tags"`
	Title       string `gorm:"index;comment:标题" json:"title"`
	Description string `gorm:"size:512;comment:描述" json:"description"`
	Content     string `gorm:"type:text;comment:内容" json:"content"`
	View        int    `gorm:"comment:点击数" json:"view"`
	Status      int    `gorm:"default:1;comment:状态 1:正常 0:禁用" json:"status"`
}
