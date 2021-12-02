package connect

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

func init() {
	//Tianjin12345!
	setPasswd := redis.DialPassword("tianjin12345")
	db := redis.DialDatabase(14)

	pool = &redis.Pool{
		MaxIdle:     8,
		MaxActive:   10,
		IdleTimeout: 100,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "172.16.52.168:6379", setPasswd, db)
		},
	}

}

func GetRedisClient() redis.Conn {
	return pool.Get()
}

func Userlogin(username, uuid string) {

	c := GetRedisClient()
	defer c.Close()

	_, err := c.Do("hset", username, "code", 1)
	if err != nil {
		fmt.Println("写入redies成功")
	} else {
		fmt.Println(err)
	}
}

func Userlogout(username string) {

	c := GetRedisClient()
	defer c.Close()

	_, err := c.Do("hset", username, 0)
	if err != nil {
		fmt.Printf("写入redies成功")
	}
}

//GetUserRedisCode 从redis读取到用户登录状态
func GetUserCode(username string) int {

	c := GetRedisClient()
	defer c.Close()

	if info, _ := redis.Int(c.Do("hget", username, "code")); string(rune(info)) == "" {
		return -1
	} else {
		return info
	}
}
func SetUserToken(username, token string) error {

	c := GetRedisClient()
	defer c.Close()

	if _, err := redis.String(c.Do("hset", username, token)); err != nil {
		return err
	} else {
		return nil
	}
}

func GetUserToken(username string) string {

	c := GetRedisClient()
	defer c.Close()
	token, _ := redis.String(c.Do("hget", username, "token"))
	if token == "" {
		return ""
	}
	return token

}

//从redis中得到uuid
func GetUseruuid(username string) string {

	c := GetRedisClient()
	defer c.Close()
	uuid, _ := redis.String(c.Do("hget", username, "uuid"))
	//返回-1代表用户未在redis中
	return uuid

}
