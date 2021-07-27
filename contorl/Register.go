package contorl

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type User struct {
	id   string
	name string
}

func Register(c *gin.Context) {

	uid, uidBool := c.GetPostForm("id")
	uname, unameBool := c.GetPostForm("name")

	if !uidBool || !unameBool {
		ErrorResponse(c)
		return
	}

	u1 := User{uid, uname}
	c.JSON(200, gin.H{
		"code":    "200",
		"message": "成功",
		"id":      u1.id,
		"name":    u1.name,
	})

	log.Printf("此次登录用户id为:%v,姓名为:%v ", u1.id, u1.name)
}

func ErrorResponse(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"code":    "400",
		"message": "id或name不能为空",
	})
}
