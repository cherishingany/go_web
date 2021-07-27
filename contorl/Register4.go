package contorl

import (
	"github.com/gin-gonic/gin"
	"go_web/connect"
	"go_web/middleware"
	userstruct "go_web/model"
)

func Register4(c *gin.Context) {

	var user userstruct.Users
	//用户结构体数据绑定
	if err := middleware.InitializationUser(&user, c); err != nil {
		return
	}
	//1.连接数据库查询用户姓名是否存在
	db := connect.InitDB()

	//
	if middleware.ExistingUser(user.Name, db) {
		c.JSON(400, gin.H{
			"code":    "400",
			"message": "用户存在",
		})
		return
	}

	//db.First(&users, "name = ?", users.Name) // 查询code为l1212的product

	//2.如不存在，创建用户，返回注册成功
	//user.UserPasswordBcrypt(c)
	//创建用户
	db.Create(&user)

	c.JSON(200, gin.H{
		"code":    200,
		"message": "注册用户成功",
	})

}
