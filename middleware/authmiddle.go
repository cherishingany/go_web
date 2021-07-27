package middleware

import (
	"github.com/gin-gonic/gin"
	"go_web/connect"
	"go_web/model"
)

func TokenAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {

		var user model.LoginUsers

		//得到Token中的username
		if GetHeaderToken(c) != "" {
			user.Name = GetHeaderToken(c)
			//去redis中查询登录状态
			//如果未登录，验证失败，断开请求
			if num := connect.GetUserCode(user.Name); num == -1 || num == 0 {
				Responce(401, "用户未登录", c)
				c.Abort()
				return
			}
			c.Next()
		}
		c.Abort()
		return

	}
}
