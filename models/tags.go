package models

type Tag struct {
	Model

	Name        string `gorm:"size:100;comment:名称" json:"name"`
	Flag        string `gorm:"index;size:50;comment:标识" json:"flag"`
	Avatar      string `gorm:"size:512;comment:图标" json:"avatar"`
	Description string `gorm:"size:512;comment:描述" json:"description"`
	Status      int    `gorm:"default:1;comment:状态 1:正常 0:禁用" json:"status"`
}

func GetTags(offset, limit int, maps map[string]interface{}) ([]*Tag, int64, error) {
	var tags []*Tag
	tagDb := db.Model(&Tag{})
	for query, args := range maps {
		tagDb.Where(query, args)
	}
	err := tagDb.Offset(offset).Limit(limit).Find(&tags).Error
	if err != nil {
		return nil, 0, err
	}
	var count int64
	err = tagDb.Limit(-1).Offset(-1).Count(&count).Error
	if err != nil {
		return nil, 0, err
	}

	return tags, count, nil
}

func GetTagByFlag(flag string) (*Tag, error) {
	var tag Tag
	err := db.Where("flag = ?", flag).First(&tag).Error
	if err != nil {
		return nil, err
	}
	return &tag, nil
}

func GetTagById(id int) (*Tag, error) {
	var tag Tag
	err := db.Where("id = ?", id).First(&tag).Error
	if err != nil {
		return nil, err
	}
	return &tag, nil
}

func AddTag(data map[string]interface{}) (*Tag, error) {
	tag := Tag{
		Name:        data["name"].(string),
		Flag:        data["flag"].(string),
		Avatar:      data["avatar"].(string),
		Description: data["description"].(string),
		Status:      data["status"].(int),
	}
	err := db.Create(&tag).Error
	if err != nil {
		return nil, err
	}
	return &tag, nil
}

func UpdateTag(id int, data map[string]interface{}) (*Tag, error) {
	tag := Tag{
		Name:        data["name"].(string),
		Flag:        data["flag"].(string),
		Avatar:      data["avatar"].(string),
		Description: data["description"].(string),
		Status:      data["status"].(int),
	}
	err := db.Model(&Tag{}).Select("name", "flag", "avatar", "description", "status", "updated_at").Where("id = ?", id).Updates(tag).Error
	if err != nil {
		return nil, err
	}
	return &tag, nil
}

func DeleteTag(id int) (bool, error) {
	tag := Tag{
		Model: Model{
			ID: id,
		},
	}
	err := db.Delete(&tag).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
