package models

type Article struct {
	Model

	Tags []Tag `gorm:"many2many:article_tags;references:" json:"tags"`

	Title       string `gorm:"size:255;index;comment:标题" json:"title"`
	Description string `gorm:"size:512;comment:描述" json:"description"`
	Content     string `gorm:"type:text;comment:内容" json:"content"`
	View        int    `gorm:"comment:点击数" json:"view"`
	Status      int    `gorm:"default:1;comment:状态 1:正常 0:禁用" json:"status"`
}

func AddArticle(data map[string]interface{}) (*Article, error) {
	article := Article{
		Tags: data["tags"].([]Tag),

		Title:       data["title"].(string),
		Description: data["description"].(string),
		Content:     data["content"].(string),
		Status:      data["status"].(int),
	}

	err := db.Create(&article).Error
	if err != nil {
		return nil, err
	}
	return &article, nil
}
