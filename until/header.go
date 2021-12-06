package until

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// postman 中间件
//http.StatusInternalServerError
func Header() gin.HandlerFunc {
	return func(c *gin.Context) {
		aggent := c.GetHeader("User-Agent")
		if aggent == "PostmanRuntime/7.26.8" {
			c.JSON(http.StatusInternalServerError, "postman")
			c.Abort()
		}

	}
}
