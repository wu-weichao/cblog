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

// @Tags 文章标签
// @Summary 获取标签列表
// @Accept  json
// @Produce  json
// @Param   name     query    string     false    "标签名称"
// @Param   flag     query    string     false    "标签flag"
// @Param   status     query    int     false    "标签状态"
// @Param   page     query    int     false    "页数"
// @Param   page_size     query    int     false    "分页条数"
// @Success 200 {string} json "{"code":200,"msg":"","data":[]}}"
// @Failure 400 {string} json "{"code":400,"msg":"","data":""}}"
// @Router /api/v1/tags [get]
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

// @Tags 文章标签
// @Summary 获取标签详情
// @Description 通过标签ID获取标签信息
// @Accept  json
// @Produce  json
// @Param   tag_id     path    int     true    "标签 ID"
// @Success 200 {string} json "{"code":200,"msg":"","data":{}}"
// @Failure 400 {string} json "{"code":400,"msg":"","data":""}}"
// @Router /api/v1/tags/{tag_id} [get]
func GetTag(c *gin.Context) {
	controller := controllers.Controller{C: c}

	id := c.Param("id")
	tagService := service.Tag{
		ID: com.StrTo(id).MustInt(),
	}
	tag, err := tagService.GetById()
	if err != nil {
		controller.Error(http.StatusBadRequest, "标签不存在", nil)
		return
	}

	controller.Success(tag, "")
}

type TagForm struct {
	Name        string `form:"name" json:"name" binding:"required,lte=100"`
	Flag        string `form:"flag" json:"flag" binding:"required,lte=50"`
	Avatar      string `form:"avatar" json:"avatar"`
	Description string `form:"description" json:"description" binding:"lte=500"`
	Status      int    `form:"status" json:"status" binding:"oneof=0 1"`
}

// @Tags 文章标签
// @Summary 新增标签
// @Accept  json
// @Produce  json
// @Param   name     formData    string     true    "标签名称"
// @Param   flag     formData    string     true    "标签标识"
// @Param   avatar     formData    string     false    "标签图标"
// @Param   description     formData    string     false    "标签描述"
// @Param   status     formData    int     false    "标签状态 1:正常 0:禁用"
// @Success 200 {string} json "{"code":200,"msg":"","data":{}}"
// @Failure 400 {string} json "{"code":400,"msg":"","data":""}}"
// @Router /api/v1/tags [post]
func CreateTag(c *gin.Context) {
	controller := controllers.Controller{C: c}

	var tagForm TagForm
	if err := c.ShouldBind(&tagForm); err != nil {
		controller.Error(http.StatusBadRequest, err.Error(), nil)
		return
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

	controller.Success(tag, "")
}

// @Tags 文章标签
// @Summary 修改标签
// @Description 通过标签ID修改标签
// @Accept  json
// @Produce  json
// @Param   tag_id     path    int     true    "标签 ID"
// @Param   name     formData    string     true    "标签名称"
// @Param   flag     formData    string     true    "标签标识"
// @Param   avatar     formData    string     false    "标签图标"
// @Param   description     formData    string     false    "标签描述"
// @Param   status     formData    int     false    "标签状态 1:正常 0:禁用"
// @Success 200 {string} json "{"code":200,"msg":"","data":{}}"
// @Failure 400 {string} json "{"code":400,"msg":"","data":""}}"
// @Router /api/v1/tags/{tag_id} [put]
func UpdateTag(c *gin.Context) {
	controller := controllers.Controller{C: c}

	var tagForm TagForm
	if err := c.ShouldBind(&tagForm); err != nil {
		controller.Error(http.StatusBadRequest, err.Error(), nil)
		return
	}

	id := com.StrTo(c.Param("id")).MustInt()
	tagService := service.Tag{
		ID:          id,
		Name:        tagForm.Name,
		Flag:        tagForm.Flag,
		Avatar:      tagForm.Avatar,
		Description: tagForm.Description,
		Status:      tagForm.Status,
	}
	// 判断记录是否存在
	existTag, _ := tagService.GetById()
	if existTag == nil {
		controller.Error(http.StatusBadRequest, "标签不存在", nil)
		return
	}
	// 判断标签是否重复
	if tagForm.Flag != existTag.Flag {
		flagTag, _ := tagService.GetByFlag()
		if flagTag != nil {
			controller.Error(http.StatusBadRequest, "标签已存在", nil)
			return
		}
	}
	// 更新
	tag, err := tagService.Update()
	if err != nil {
		controller.Error(http.StatusBadRequest, "更新成功", nil)
		return
	}
	controller.Success(tag, "")
}

// @Tags 文章标签
// @Summary 删除标签
// @Description 通过标签ID删除标签
// @Accept  json
// @Produce  json
// @Param   tag_id     path    int     true    "标签 ID"
// @Success 200 {string} json "{"code":200,"msg":"","data":""}"
// @Failure 400 {string} json "{"code":400,"msg":"","data":""}}"
// @Router /api/v1/tags/{tag_id} [delete]
func DeleteTag(c *gin.Context) {
	controller := controllers.Controller{C: c}
	id := com.StrTo(c.Param("id")).MustInt()
	tagService := service.Tag{
		ID: id,
	}
	_, err := tagService.GetById()
	if err != nil {
		controller.Error(http.StatusBadRequest, "标签不存在", nil)
		return
	}
	result, err := tagService.Delete()
	if err != nil || result == false {
		controller.Error(http.StatusBadRequest, "标签删除失败", nil)
		return
	}

	controller.Success(id, "操作成功")
}
