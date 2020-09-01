package middleware

import (
	"cblog/controllers"
	"cblog/models"
	"cblog/pkg/setting"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// jwt 校验中间件
func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		flag := true
		errMsg := "Token鉴权失败"
		token := c.GetHeader("Authorization")
		if token == "" {
			flag = false
		} else {
			token = token[7:]
			// 检验Token
			_, err := (&Jwt{}).ParseToken(token)
			if err != nil {
				flag = false
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorMalformed:
					errMsg = "Token格式错误"
				case jwt.ValidationErrorExpired:
					errMsg = "Token已过期"
				case jwt.ValidationErrorNotValidYet:
					errMsg = "Token验证错误"
				}
			}
		}

		if flag == false {
			c.JSON(http.StatusUnauthorized, controllers.Response{
				Code:    http.StatusUnauthorized,
				Message: errMsg,
				Data:    "",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

type CustomClaims struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	Email  string `json:"email"`
	Phone  string `json:"phone"`
	Type   int    `json:"type"`
	jwt.StandardClaims
}

type Jwt struct {
	JwtSecret []byte
}

func (this *Jwt) CreateToken(user *models.User) (string, int64, error) {
	this.JwtSecret = []byte(setting.AppSetting.JwtSecret)

	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(setting.AppSetting.JwtTtl) * time.Hour)

	claims := CustomClaims{
		ID:     user.ID,
		Name:   user.Name,
		Avatar: user.Avatar,
		Email:  user.Email,
		Phone:  user.Phone,
		Type:   user.Type,
		StandardClaims: jwt.StandardClaims{
			NotBefore: nowTime.Unix() - 1000,
			ExpiresAt: expireTime.Unix(),
			Issuer:    "cblog",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(this.JwtSecret)

	return token, expireTime.Unix(), err
}

func (this *Jwt) ParseToken(tokenString string) (*CustomClaims, error) {
	this.JwtSecret = []byte(setting.AppSetting.JwtSecret)

	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return this.JwtSecret, nil
	})

	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
	}
	return nil, err

}
