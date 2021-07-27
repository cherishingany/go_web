package contorl

import (
	"github.com/gin-gonic/gin"
	"go_web/middleware"
	"go_web/service"
)

func CreateVideo(c *gin.Context) {

	var newVideo service.CreateVideo
	if err := middleware.InitializationUser(&newVideo, c); err == nil {
		service.NewVideo(&newVideo)
	} else {
		c.JSON(200, "创建失败")
	}

}
