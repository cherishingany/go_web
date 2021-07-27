package service

import (
	"github.com/gin-gonic/gin"
	"go_web/middleware"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginUser struct {
	Name     string `json:"name"  binding:"required,gt=3,lt=6" gorm:"not null" ` // 10 <= age <= 120。
	Password string `json:"password"  binding:"required,gt=4,lt=10"`
}

//用户密码判断
func (user *LoginUser) UserPasswordLogin(c *gin.Context, db *gorm.DB, password string) bool {
	db.Where("name = ?", user.Name).First(&user)
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err == nil {
		return true
	}
	middleware.Responce(200, "用户名密码错误", c)
	return false
}
