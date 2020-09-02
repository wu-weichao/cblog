package service

import (
	"cblog/models"
	"strconv"
	"strings"
)

type Article struct {
	ID          int
	Tags        string
	Title       string
	Description string
	Content     string
	View        int
	Status      int
}

func (this *Article) GetById() (*models.Article, error) {
	return models.GetArticleById(this.ID)
}

func (this *Article) GetAll(offset, limit int) (articles []*models.Article, count int64, err error) {
	maps := make(map[string]interface{})

	if this.Tags != "" {
		var tags []int
		for _, tagId := range strings.Split(this.Tags, ",") {
			tempId, _ := strconv.Atoi(tagId)
			tags = append(tags, tempId)
		}
		maps["tags"] = tags
	}
	if this.Status != -1 {
		maps["status = ?"] = this.Status
	}

	articles, err = models.GetArticles(offset, limit, maps)
	if err != nil {
		return
	}
	//count, err = models.GetArticlesCount(maps)
	//if err != nil {
	//	return
	//}

	return
}

func (this *Article) Add() (*models.Article, error) {
	var tags []models.Tag
	for _, tagId := range strings.Split(this.Tags, ",") {
		tempId, _ := strconv.Atoi(tagId)
		tags = append(tags, models.Tag{Model: models.Model{ID: tempId}})
	}

	data := map[string]interface{}{
		"tags":        tags,
		"title":       this.Title,
		"description": this.Description,
		"content":     this.Content,
		"status":      this.Status,
	}
	return models.AddArticle(data)
}
