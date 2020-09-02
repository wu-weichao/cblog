package v1

import (
	"cblog/controllers"
	"cblog/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetArticles(c *gin.Context) {

}

func GetArticle(c *gin.Context) {

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
	}

	articleService := service.ArticleService{
		Tags:        articleForm.Tags,
		Title:       articleForm.Title,
		Description: articleForm.Description,
		Content:     articleForm.Content,
		Status:      articleForm.Status,
	}
	article, err := articleService.Add()
	if err != nil {
		controller.Error(http.StatusBadRequest, err.Error(), nil)
	}

	controller.Success(article, "")
}

func UpdateArticle(c *gin.Context) {

}

func DeleteArticle(c *gin.Context) {

}
