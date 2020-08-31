package v1

import (
	"cblog/controllers"
	"cblog/models"
	"cblog/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

func GetTags(c *gin.Context) {
	controller := controllers.Controller{C: c}
	controller.SetPage()

	name := c.Query("name")
	flag := c.Query("flag")
	status := -1
	if arg := c.Query("status"); arg != "" {
		status = com.StrTo(arg).MustInt()
	}
	tagService := service.Tag{
		Name:   name,
		Flag:   flag,
		Status: status,
	}

	tags, total, err := tagService.GetAll(controller.Offset, controller.Limit)
	if err != nil {
		fmt.Printf("GetTags Error: %v", err)
		controller.Error(http.StatusBadRequest, "查询失败，请稍后再试", nil)
		return
	}

	controller.Paginate(tags, total)
}

func GetTag(c *gin.Context) {
	controller := controllers.Controller{C: c}

	id := c.Param("id")
	tagService := service.Tag{
		Id: com.StrTo(id).MustInt(),
	}
	tag, err := tagService.GetById()
	if err != nil {
		controller.Error(http.StatusBadRequest, "标签不存在", nil)
		return
	}

	controller.Success("", tag)
}

type TagForm struct {
	Name        string `form:"name" json:"name" binding:"required,lte=100"`
	Flag        string `form:"flag" json:"flag" binding:"required,lte=50"`
	Avatar      string `form:"avatar" json:"avatar"`
	Description string `form:"description" json:"description" binding:"lte=500"`
	Status      int    `form:"status" json:"status" binding:"oneof=0 1"`
}

func CreateTag(c *gin.Context) {
	controller := controllers.Controller{C: c}

	var tagForm TagForm
	if err := c.ShouldBind(&tagForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	tagService := service.Tag{
		Name:        tagForm.Name,
		Flag:        tagForm.Flag,
		Avatar:      tagForm.Avatar,
		Description: tagForm.Description,
		Status:      tagForm.Status,
	}
	_, err := tagService.GetByFlag()
	if err == nil {
		controller.Error(http.StatusBadRequest, "标签已存在", nil)
		return
	}
	// add tag
	var tag *models.Tag
	tag, err = tagService.Add()
	if err != nil {
		controller.Error(http.StatusBadRequest, "标签添加失败", nil)
		return
	}

	controller.Success("标签添加成功", tag)
	//c.JSON(http.StatusOK, gin.H{"error": "标签添加失败"})
}

func UpdateTag(c *gin.Context) {

}

func DeleteTag(c *gin.Context) {

}
