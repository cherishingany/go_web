package contorl

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_web/connect"
	"go_web/middleware"
	"go_web/model"
	"golang.org/x/crypto/bcrypt"
)

//解决依赖问题, 完成
// "error": {}, 完成
//映射标签问题, 完成
//增加删除接口，增加id自增和name主键,完成

//Login 用户登录接口
func Login(c *gin.Context) {

	//1.绑定结构体，到redis读取用户登录状态，如果已经登录，则登录失败。else发放进行token
	var user model.LoginUsers
	//获取json数据
	if err := middleware.InitializationUser(&user, c); err != nil {
		return
	}

	if code := connect.GetUserCode(user.Name); code == 1 {
		middleware.Responce(401, "用户已经登录", c)
		return
	}

	//用户未登录，得到他的用户名密码和mysql中进行对比，如果相同则登录成功

	userpassword := user.Password
	db := connect.InitDB()

	// 验证token用户是否存在
	if !middleware.ExistingUserScan(&user, db) {
		middleware.Responce(401, "token用户不存在", c)
		return
	}

	//连接数据库，查询用户是否与上下文的用户密码一致
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userpassword)); err != nil {
		middleware.Responce(401, "用户名密码错误", c)
		fmt.Println(userpassword)
		fmt.Println(user.Password)
		return
	}
	//发放token
	if err := middleware.UserToken(&user, c); err != nil {
		return
	}
	connect.Userlogin(user.Name, "124324234")

}
