package contorl

import (
	"github.com/gin-gonic/gin"
	"go_web/connect"
)

func DeleteUser(c *gin.Context) {

	//var user model.Users
	var user interface{}
	//
	//if err := middleware.InitializationUser(&user, c); err != nil {
	//	return
	//}

	db := connect.InitDB()
	info := db.Table("users").First(&user, "name = ?", "BeJ5")

	print(info.RowsAffected)
	//db.Debug().Raw("select name, password from users where name = ? and deleted_at is null", user.Name)
	//users := find()
	//for _, user = range users {
	//
	//}

	c.JSON(200, gin.H{
		"code":    "200",
		"message": user,
	})
	return

}

//func find() []model.Users {
//	users := make([]model.Users, 10)
//
//	db := connect.InitDB()
//
//	//db.Find(&user)
//	//db.Debug().Find(&users)
//	db.Debug().Table("users").Select("name,password").Where("name = ?", username)
//	return users
//}
