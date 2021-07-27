package service

import (
	"github.com/gin-gonic/gin"
	"go_web/connect"
	"go_web/middleware"
	"go_web/model"
)

func UserLogout(logoutuser *model.LoginUsers, c *gin.Context) {

	if num := connect.GetUserCode(logoutuser.Name); num == 0 || num == -1 {
		middleware.Responce(401, "用户未登录", c)
		c.Abort()
		return
	}
	connect.Userlogout(logoutuser.Name)
	middleware.Responce(500, "成功登出", c)
}
