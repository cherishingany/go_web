package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go_web/connect"
	userstruct "go_web/model"
	"time"
)

type Key struct {
	UserName string `json:"user_name"`
	Time     string `json:"local_time"`
	jwt.StandardClaims
}

var Secret = []byte("User_Secret")

//发放token
func UserToken(user *userstruct.LoginUsers, c *gin.Context) error {

	key := &Key{
		UserName: user.Name,
		Time:     time.Now().Format("2006-01-02 15:04:05"),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(10 * time.Hour).Unix(),
			Subject:   "user_token",
			Issuer:    "oceanlearn.tech",
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, key)
	logintoken, err := token.SignedString(Secret)
	if err != nil {
		c.JSON(500, gin.H{
			"code":    "500",
			"message": "系统错误，生成token失败",
			"err":     err.Error(),
		})
		return err
	}

	//返回数据
	c.JSON(200, gin.H{
		"code":    "200",
		"message": "登录成功",
		"data": gin.H{
			"token": logintoken,
		},
	})

	_ = connect.SetUserToken(user.Name, logintoken)

	return nil

}

func ParseToken(tokenstring string) (*jwt.Token, *Key, error) {
	claims := &Key{}
	//var name Key//这里不理解

	token, err := jwt.ParseWithClaims(tokenstring, claims, func(token *jwt.Token) (interface{}, error) {
		return Secret, nil
	})

	return token, claims, err

}
