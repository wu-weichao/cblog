package v1

import (
	"cblog/controllers"
	"cblog/service"
	"github.com/gin-gonic/gin"
	//"cblog/middleware"
	"net/http"
)

func Register(c *gin.Context) {

}

type LoginForm struct {
	Email    string `form:"email" json:"email" binding:"required,lte=255"`
	Password string `form:"password" json:"password" binding:"required,lte=255"`
}

func Login(c *gin.Context) {
	controller := controllers.Controller{C: c}

	var loginForm LoginForm
	if err := c.ShouldBind(&loginForm); err != nil {
		controller.Error(http.StatusBadRequest, err.Error(), nil)
		return
	}

	var userService = service.User{
		Email:    loginForm.Email,
		Password: loginForm.Password,
	}
	token, expiresAt, user := userService.Login()
	if token == "" {
		controller.Error(http.StatusBadRequest, "账号或密码错误", nil)
		return
	}
	response := map[string]interface{}{
		"user":        user,
		"token":       token,
		"expire_time": expiresAt,
	}

	controller.Success(response, "")
}

func Logout(c *gin.Context) {

}

func GetUser(c *gin.Context) {
	userId := c.Param("id")

	c.JSON(200, gin.H{"user_id": userId})
}
