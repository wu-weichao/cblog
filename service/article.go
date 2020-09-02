package service

import (
	"cblog/models"
	"strconv"
	"strings"
)

type ArticleService struct {
	ID          int
	Tags        string
	Title       string
	Description string
	Content     string
	View        int
	Status      int
}

func (this *ArticleService) Add() (*models.Article, error) {
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
