package middleware

import (
	"github.com/gin-gonic/gin"
	"go_web/model"
	"gorm.io/gorm"
	"log"
	"strings"
)

//InitializationUser 结构体绑定
func InitializationUser(user interface{}, c *gin.Context) (err error) {
	if err := c.ShouldBindJSON(user); err != nil {
		return err
	}
	return nil
}

//ExistingUser 查询用户是否存在
func ExistingUser(username string, db *gorm.DB) bool {
	//db.Where("name = ?", user.Name).First(&user)
	if result := db.Where("name = ?", username).First(&model.Users{}); result.RowsAffected != 0 {

		//db.Debug().Model(&Users{}).Find(&SelectUsers{}).Where("name = ?", user.Name)
		//db.Debug().Select("name", "password").Model(&Users{}).Where("name = ?", user.Name).Scan(user)
		//db.Debug().Raw("select name, password from users where name = ? and deleted_at is null;
		log.Printf("用户存在，共找到 %v 条数据 ", result.RowsAffected)
		return true
	} else {
		log.Println(username)
		log.Println(result.Error)
		return false
	}

}

func ExistingUserScan(user *model.LoginUsers, db *gorm.DB) bool {
	//db.Where("name = ?", user.Name).First(&user)
	if result := db.Debug().Raw("select name, password from users where name = ? and deleted_at is null;", user.Name).Scan(&user); result.RowsAffected != 0 {
		//db.Where("name = ?", user.Name).First(&model.Users{}).Scan(&user)
		//db.Debug().Model(&Users{}).Find(&SelectUsers{}).Where("name = ?", user.Name)
		//db.Debug().Select("name", "password").Model(&Users{}).Where("name = ?", user.Name).Scan(user)
		//db.Debug().Raw("select name, password from users where name = ? and deleted_at is null;
		log.Printf("用户存在，共找到 %v 条数据 ", result.RowsAffected)
		return true
	} else {
		log.Println(result.Error)
		return false
	}

}

//GetHeaderToken 从token中取到username
func GetHeaderToken(c *gin.Context) string {
	tokenstring := c.GetHeader("Authentication")
	if tokenstring == "" || !strings.HasPrefix(tokenstring, "Bearer ") {
		Responce(401, "权限不足", c)
		return ""
	}
	tokenstring = tokenstring[7:] //得到header中的token
	token, claims, err := ParseToken(tokenstring)

	if err != nil || !token.Valid {
		Responce(401, "InvalidToken", c)
		log.Println(err)
		return ""
	}
	return claims.UserName
}
func Responce(code int, message interface{}, c *gin.Context) {
	c.JSON(200, gin.H{
		"code":    code,
		"message": message,
	})
}
