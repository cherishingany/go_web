package engine

import (
	"github.com/gin-gonic/gin"
	"go_web/contorl"
)

func NewEngine() *gin.Engine {

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "你好，gin",
		})
	})

	user := router.Group("/user") //创建路由组
	//user.Use(middleware.Header())
	{

		user.POST("/register3", contorl.Register3)
		user.POST("/login", contorl.Login)

		login := user.Group("")
		//login.Use(middleware.TokenAuthentication())

		{
			login.POST("/info", contorl.Info)
			login.POST("/logout", contorl.Logout)
			login.POST("/createvideo", contorl.CreateVideo)
			login.GET("/:name", contorl.VideoList)
		}

		user.POST("/deleteuser", contorl.DeleteUser)

	}

	//r.GET("/error", func(c *gin.Context) {
	//	c.HTML(404, "error.html", gin.H{
	//		"name": "haha",
	//	})
	//})
	//
	//r.GET("/indexlogin", func(c *gin.Context) {
	//	c.HTML(200, "index1.html",nil)
	//})

	return router
}
