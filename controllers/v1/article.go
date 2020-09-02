package v1

import (
	"cblog/controllers"
	"cblog/service"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

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

func UpdateArticle(c *gin.Context) {

}

func DeleteArticle(c *gin.Context) {

}
