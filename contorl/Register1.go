package contorl

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"log"
	"net/http"
)

//表单实体绑定
//可调用http.StatusOK来调用系统系统的返回值
//可调用http.StatusText(203)返回系统文本
//加载html文件

func Register1(c *gin.Context) {
	var s1 Student
	if err := c.ShouldBindBodyWith(&s1, binding.JSON); err != nil {
		c.JSON(400, gin.H{
			"code":    "400",
			"message": "结构体初始化失败",
		})
	} else {
		res := http.StatusText(200)
		c.JSON(http.StatusOK, gin.H{
			"code":    "200",
			"message": res,
		})

		log.Printf("用户的姓名为: %v,年龄为: %v,地址为: %v", s1.Name, s1.Age, s1.Address)

	}

}

type Student struct {
	Num     int    `form:"age" bind:"required"`
	Name    string `form:"name" bind:"required"` //必须添加属性字段
	Age     int    `form:"age" bind:"required"`
	Address string `form:"address" bind:"required"`
}

//可以把code message赋值为一个结构体，然后初始化结构体并返回
type ErrResonance struct {
	code    int
	message string
}
