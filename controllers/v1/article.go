package v1

import (
	"cblog/controllers"
	"cblog/service"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

// @Tags 文章
// @Summary 获取文章列表
// @Version v1
// @Accept  json
// @Produce  json
// @Param   tag     query    string     false    "标签ID,多个以逗号分隔"
// @Param   status     query    int     false    "文章状态"
// @Param   page     query    int     false    "页数"
// @Param   page_size     query    int     false    "分页条数"
// @Success 200 {string} json "{"code":200,"msg":"","data":[]}"
// @Failure 400 {string} json "{"code":400,"msg":"","data":""}}"
// @Router /api/v1/articles [get]
func GetArticles(c *gin.Context) {
	controller := controllers.Controller{C: c}
	controller.SetPage()

	tags := c.Query("tag")
	status := -1
	if arg := c.Query("status"); arg != "" {
		status = com.StrTo(arg).MustInt()
	}
	articleService := service.Article{
		Tags:   tags,
		Status: status,
	}

	articles, total, err := articleService.GetAll(controller.Offset, controller.Limit)
	if err != nil {
		controller.Error(http.StatusBadRequest, "查询失败，请稍后再试", nil)
		return
	}

	controller.Paginate(articles, total)
}

// @Tags 文章
// @Summary 获取文章详情
// @Description 通过文章ID获取文章详情
// @Version v1
// @Accept  json
// @Produce  json
// @Param   article_id     path    int     true    "文章 ID"
// @Success 200 {string} json "{"code":200,"msg":"","data":{}}"
// @Failure 400 {string} json "{"code":400,"msg":"","data":""}}"
// @Router /api/v1/articles/{articel_id} [get]
func GetArticle(c *gin.Context) {
	controller := controllers.Controller{C: c}

	id := com.StrTo(c.Param("id")).MustInt()
	articleService := service.Article{
		ID: id,
	}
	article, err := articleService.GetById()
	if err != nil {
		controller.Error(http.StatusBadRequest, err.Error(), nil)
		return
	}

	controller.Success(article, "")
}

type ArticleForm struct {
	Tags        string `form:"tags" json:"tags" binding:"required"`
	Title       string `form:"title" json:"title" binding:"required,lte=200"`
	Description string `form:"description" json:"description" binding:"lte=500"`
	Content     string `form:"content" json:"content" binding:"required"`
	Status      int    `form:"status" json:"status" binding:"oneof=0 1"`
}

// @Tags 文章
// @Summary 新增文章
// @Version v1
// @Accept  json
// @Produce  json
// @Param   tags     formData    string     true    "标签 ID"
// @Param   title     formData    string     true    "标题"
// @Param   description     formData    string     false    "描述"
// @Param   content     formData    string     true    "内容"
// @Param   status     formData    int     false    "状态 1:正常 0:禁用"
// @Success 200 {string} json "{"code":200,"msg":"","data":{}}"
// @Failure 400 {string} json "{"code":400,"msg":"","data":""}}"
// @Router /api/v1/articles [post]
func CreateArticle(c *gin.Context) {
	controller := controllers.Controller{C: c}

	var articleForm ArticleForm
	if err := c.ShouldBind(&articleForm); err != nil {
		controller.Error(http.StatusBadRequest, err.Error(), nil)
		return
	}
	// 判断Tags对应的ID是否存在
	if (&service.Tag{}).CheckIdExists(articleForm.Tags) == false {
		controller.Error(http.StatusBadRequest, "文章标签错误或不存在", nil)
		return
	}

	articleService := service.Article{
		Tags:        articleForm.Tags,
		Title:       articleForm.Title,
		Description: articleForm.Description,
		Content:     articleForm.Content,
		Status:      articleForm.Status,
	}
	article, err := articleService.Add()
	if err != nil {
		controller.Error(http.StatusBadRequest, err.Error(), nil)
		return
	}

	controller.Success(article, "")
}

// @Tags 文章
// @Summary 修改文章
// @Description 通过文章ID修改文章
// @Version v1
// @Accept  json
// @Produce  json
// @Param   article_id     path    int     true    "文章 ID"
// @Param   tags     formData    string     true    "标签 ID"
// @Param   title     formData    string     true    "标题"
// @Param   description     formData    string     false    "描述"
// @Param   content     formData    string     true    "内容"
// @Param   status     formData    int     false    "状态 1:正常 0:禁用"
// @Success 200 {string} json "{"code":200,"msg":"","data":{}}"
// @Failure 400 {string} json "{"code":400,"msg":"","data":""}}"
// @Router /api/v1/articles/{articel_id} [put]
func UpdateArticle(c *gin.Context) {
	controller := controllers.Controller{C: c}

	id := com.StrTo(c.Param("id")).MustInt()

	var articleForm ArticleForm
	if err := c.ShouldBind(&articleForm); err != nil {
		controller.Error(http.StatusBadRequest, err.Error(), nil)
		return
	}

	articleService := service.Article{
		ID:          id,
		Tags:        articleForm.Tags,
		Title:       articleForm.Title,
		Description: articleForm.Description,
		Content:     articleForm.Content,
		Status:      articleForm.Status,
	}
	_, err := articleService.GetById()
	if err != nil {
		controller.Error(http.StatusBadRequest, "文章不存在", nil)
		return
	}
	// 判断Tags对应的ID是否存在
	if (&service.Tag{}).CheckIdExists(articleForm.Tags) == false {
		controller.Error(http.StatusBadRequest, "文章标签错误或不存在", nil)
		return
	}
	article, err := articleService.Update()
	if err != nil {
		controller.Error(http.StatusBadRequest, err.Error(), nil)
		return
	}

	controller.Success(article, "")
}

// @Tags 文章
// @Summary 删除文章
// @Description 通过文章ID删除文章
// @Version v1
// @Accept  json
// @Produce  json
// @Param   article_id     path    int     true    "文章 ID"
// @Success 200 {string} json "{"code":200,"msg":"","data":""}}"
// @Failure 400 {string} json "{"code":400,"msg":"","data":""}}"
// @Router /api/v1/articles/{articel_id} [delete]
func DeleteArticle(c *gin.Context) {
	controller := controllers.Controller{C: c}

	id := com.StrTo(c.Param("id")).MustInt()
	articleService := service.Article{
		ID: id,
	}
	result, err := articleService.Delete()
	if err != nil || result == false {
		controller.Error(http.StatusBadRequest, "文章删除失败", nil)
		return
	}

	controller.Success(id, "")
}
