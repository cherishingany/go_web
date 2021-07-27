package contorl

import (
	"github.com/gin-gonic/gin"
	"go_web/middleware"
	"go_web/model"
	"go_web/service"
)

func Logout(c *gin.Context) {
	var logoutuser model.LoginUsers

	if err := middleware.InitializationUser(&logoutuser, c); err != nil {
		return
	}
	service.UserLogout(&logoutuser, c)

}
