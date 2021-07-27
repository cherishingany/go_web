package contorl

import (
	"github.com/gin-gonic/gin"
	"go_web/middleware"
	"go_web/model"
)

type UserLoginService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=40"`
	Token    bool   `form:"token" json:"token"`
}

type UserInfo struct {
	Name  string
	Token string `form:"token" json:"token ,omitempty"`
}

func Info(c *gin.Context) {

	var user model.Users

	if err := middleware.InitializationUser(&user, c); err != nil {
		return
	}
	u1 := UserInfo{
	}
	if user.Token {
		u1.Name = "lili"
		u1.Token = "123"
		c.JSON(200, u1)
		return
	} else {
		u1.Name = "lili"

		c.JSON(200, u1)
		return
	}

}
