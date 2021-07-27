package contorl

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"go_web/middleware"
	"go_web/service"
)

//----------------------------------------------------第三版-------------------------------------------------
//https://blog.csdn.net/weixin_42279809/article/details/107800081
//利用指针问题解决结构体初始化传值为nil
//结构体字段不用一个属性一个属性赋值

//defer close
//初始化结构体字段不是必须的问题 ,解决，binding写错了
//嵌入结构体
//结构体便利顺序问题 ，未

func Register3(c *gin.Context) {
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
