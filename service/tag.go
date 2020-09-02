package service

import (
	"cblog/models"
	"strings"
)

type Tag struct {
	ID          int
	Name        string
	Flag        string
	Avatar      string
	Description string
	Status      int
}

func (this *Tag) CheckIdExists(ids string) bool {
	idArr := strings.Split(ids, ",")

	maps := make(map[string]interface{})
	maps["id in ?"] = idArr
	count, _ := models.GetTagsCount(maps)

	if len(idArr) == int(count) {
		return true
	}
	return false
}

func (this *Tag) GetAll(offset, limit int) (tags []*models.Tag, count int64, err error) {
	maps := make(map[string]interface{})
	if this.Name != "" {
		maps["name LIKE ?"] = "%" + this.Name + "%"
	}
	if this.Flag != "" {
		maps["flag LIKE ?"] = "%" + this.Flag + "%"
	}
	if this.Status != -1 {
		maps["status = ?"] = this.Status
	}

	tags, err = models.GetTags(offset, limit, maps)
	if err != nil {
		return
	}
	count, err = models.GetTagsCount(maps)
	if err != nil {
		return
	}

	return
}

func (this *Tag) GetByFlag() (*models.Tag, error) {
	return models.GetTagByFlag(this.Flag)
}

func (this *Tag) GetById() (*models.Tag, error) {
	return models.GetTagById(this.ID)
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

func (this *Tag) Update() (*models.Tag, error) {
	data := map[string]interface{}{
		"name":        this.Name,
		"flag":        this.Flag,
		"avatar":      this.Avatar,
		"description": this.Description,
		"status":      this.Status,
	}
	return models.UpdateTag(this.ID, data)
}

func (this *Tag) Delete() (bool, error) {
	return models.DeleteTag(this.ID)
}
