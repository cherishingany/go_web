package until

import "github.com/gin-gonic/gin"

// postman 中间件

func Header() gin.HandlerFunc {
	return func(c *gin.Context) {
		aggent := c.GetHeader("User-Agent")
		if aggent == "PostmanRuntime/7.26.8" {
			c.JSON(200, "postman")
			c.Abort()
		}

	}
}
