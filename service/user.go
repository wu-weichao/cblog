package service

import (
	"cblog/middleware"
	"cblog/models"
)

type User struct {
	ID       int
	Name     string
	Password string
	Avatar   string
	Email    string
	Phone    string
	Type     int
	Status   int
}

func (this *User) Login() (token string, expiresAt int64, user *models.User) {
	// 获取用户
	user, err := models.GetUserByEmail(this.Email)
	// 校验密码
	if err != nil || this.Password != user.Password {
		return
	}
	// 生成token
	jwt := middleware.Jwt{}
	token, expiresAt, err = jwt.CreateToken(user)
	if err != nil {
		return
	}
	return
}
