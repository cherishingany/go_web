package service

import (
	"github.com/gin-gonic/gin"
	"go_web/connect"
	"go_web/middleware"
	"go_web/model"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type RegisterUser struct {
	Name     string `json:"name"  binding:"required,gt=3,lt=6" gorm:"not null" ` // 10 <= age <= 120。
	Password string `json:"password"  binding:"required,gt=4,lt=10"`
	Age      uint   `json:"age"  binding:"required"`
	Address  string `json:"address"  binding:"required"`
}

//用户密码加密
func (registerUser *RegisterUser) UserPasswordBcrypt(c *gin.Context) {
	Password, err := bcrypt.GenerateFromPassword([]byte(registerUser.Password), bcrypt.DefaultCost)
	if err == nil {
		registerUser.Password = string(Password)
		return
	}
	c.JSON(500, gin.H{
		"code":    500,
		"message": "系统错误，用户密码加密失败",
		"error":   err,
	})
	return
}

func Register(registerUser *RegisterUser, c *gin.Context) {
	//创建数据库连接
	db := connect.InitDB()

	//判断用户名是否存在
	if middleware.ExistingUser(registerUser.Name, db) {
		c.JSON(200, gin.H{
			"code":    "400",
			"message": "用户已经存在",
		})
		return
	}
	//用户密码加密
	registerUser.UserPasswordBcrypt(c)


	 user  := &model.Users{
		Name:     registerUser.Name,
		Password: registerUser.Password,
		Age:      registerUser.Age,
		Address:  registerUser.Address,
	}


	//创建用户
	if res := db.Create(&user); res.Error  == nil {
		middleware.Responce(200,"注册用户成功",c)
	}else {
		log.Println(res.Error)
		middleware.Responce(200,"注册用户失败",c)
	}

}
