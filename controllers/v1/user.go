package v1

import (
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {

}

func Login(c *gin.Context) {

}

func Logout(c *gin.Context) {

}

func GetUser(c *gin.Context) {
	userId := c.Param("id")

	c.JSON(200, gin.H{"user_id": userId})
}
