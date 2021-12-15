package contorl

import (
	"github.com/gin-gonic/gin"
	"go_web/middleware"
	"go_web/service"
)

func VideoList(c *gin.Context) {
	if username := c.Param("name"); username != "" {
		res := service.GetVideoList(username)
		middleware.Responce(200, res, c)
		return
	}
	middleware.Responce(500, "请求错误", c)
}
