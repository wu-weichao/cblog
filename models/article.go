package models

import (
	"gorm.io/gorm"
)

type Article struct {
	Model

	Tags []Tag `gorm:"many2many:article_tags;references:" json:"tags"`

	Title       string `gorm:"size:255;index;comment:标题" json:"title"`
	Description string `gorm:"size:512;comment:描述" json:"description"`
	Content     string `gorm:"type:text;comment:内容" json:"content"`
	View        int    `gorm:"comment:点击数" json:"view"`
	Status      int    `gorm:"default:1;comment:状态 1:正常 0:禁用" json:"status"`
}

func GetArticles(offset, limit int, maps map[string]interface{}) (articles []*Article, err error) {
	if maps["tags"] != nil {
		var articleIds []int
		result := make(map[string]interface{})
		db.Table("article_tags").Select("article_id").Where("tag_id in ?", maps["tags"]).Find(&result)

		for _, aid := range result {
			articleIds = append(articleIds, int(aid.(int64)))
		}
		maps["id in ?"] = articleIds
		delete(maps, "tags")
	}

	tagDb := db.Model(&Article{}).Preload("Tags")
	for query, args := range maps {
		tagDb.Where(query, args)
	}
	err = tagDb.Order("created_at desc").Offset(offset).Limit(limit).Find(&articles).Error
	if err != nil {
		return nil, err
	}

	return
}

func GetArticlesCount(maps map[string]interface{}) (count int64, err error) {
	if maps["tags"] != nil {
		var articleIds []int
		result := make(map[string]interface{})
		db.Table("article_tags").Select("article_id").Where("tag_id in ?", maps["tags"]).Find(&result)

		for _, aid := range result {
			articleIds = append(articleIds, int(aid.(int64)))
		}
		maps["id in ?"] = articleIds
		delete(maps, "tags")
	}

	tagDb := db.Model(&Article{}).Preload("Tags")
	for query, args := range maps {
		tagDb.Where(query, args)
	}
	err = tagDb.Count(&count).Error
	if err != nil {
		return 0, err
	}
	return
}

func GetArticleById(id int) (*Article, error) {
	var article Article
	err := db.Preload("Tags", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "name")
	}).First(&article, id).Error
	if err != nil {
		return nil, err
	}
	return &article, nil
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

func UpdateArticle(data map[string]interface{}) (*Article, error) {
	var article Article
	err := db.First(&article, data["id"]).Error
	if err != nil {
		return nil, err
	}
	tx := db.Begin()
	article.Title = data["title"].(string)
	article.Description = data["description"].(string)
	article.Content = data["content"].(string)
	article.Status = data["status"].(int)
	if err := tx.Save(&article).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	// 更新关联关系
	if err := tx.Model(&article).Association("Tags").Replace(data["tags"].([]Tag)); err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return &article, nil
}

func DeleteArticle(id int) (bool, error) {
	article := Article{
		Model: Model{
			ID: id,
		},
	}
	tx := db.Begin()
	if err := tx.Delete(&article).Error; err != nil {
		tx.Rollback()
		return false, err
	}
	if err := tx.Model(&article).Association("Tags").Clear(); err != nil {
		tx.Rollback()
		return false, err
	}
	tx.Commit()

	return true, nil
}
