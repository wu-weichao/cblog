package service

import (
	"cblog/models"
)

type Tag struct {
	Id          int
	Name        string
	Flag        string
	Avatar      string
	Description string
	Status      int
}

func (this *Tag) GetAll(offset, limit int) ([]*models.Tag, int64, error) {
	maps := make(map[string]interface{})
	if this.Name != "" {
		maps["name LIKE ?"] = "%" + this.Name + "%"
	}
	if this.Flag != "" {
		maps["flag LIKE ?"] = "%" + this.Flag + "%"
	}
	if this.Status != -1 {
		maps["status"] = this.Status
	}
	return models.GetTags(offset, limit, maps)
}

func (this *Tag) GetByFlag() (*models.Tag, error) {
	return models.GetTagByFlag(this.Flag)
}

func (this *Tag) GetById() (*models.Tag, error) {
	return models.GetTagById(this.Id)
}

func (this *Tag) Add() (*models.Tag, error) {
	data := map[string]interface{}{
		"name":        this.Name,
		"flag":        this.Flag,
		"avatar":      this.Avatar,
		"description": this.Description,
		"status":      this.Status,
	}
	return models.AddTag(data)
}
