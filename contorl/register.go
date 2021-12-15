package contorl

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"go_web/middleware"
	"go_web/service"
)

//----------------------------------------------------第三版-------------------------------------------------
//https://blog.csdn.net/weixin_42279809/article/details/107800081

func Register(c *gin.Context) {
	//用户结构体数据绑定
	var registerUser service.RegisterUser
	if err := middleware.InitializationUser(&registerUser, c); err == nil {
		service.Register(&registerUser, c)
	} else {
		c.JSON(200, gin.H{
			"code":    400,
			"message": "结构体初始化失败",
			"error":   err.Error(),
		})
	}
}
