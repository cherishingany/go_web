package contorl

//
//import (
//	"github.com/gin-gonic/gin"
//	"github.com/gin-gonic/gin/binding"
//	_ "github.com/go-sql-driver/mysql"
//	"gorm.io/driver/mysql"
//	"gorm.io/gorm"
//	"log"
//)
//
////defer close
////初始化结构体字段不是必须的问题
////嵌入结构体
////结构体便利顺序问题
//func Register2(c *gin.Context) {
//
//	var User1 Users
//
//	if err := c.ShouldBindBodyWith(&User1, binding.JSON); err != nil || User1.ID == 0 || User1.Age == 0 {
//		c.JSON(400, gin.H{
//			"code":    "400",
//			"message": "结构体初始化失败",
//			"error":   err,
//		})
//		log.Println(err)
//	}
//
//
//	log.Printf("用户id为: %v ,用户的姓名为: %v,年龄为: %v,地址为: %v", User1.ID, User1.Name, User1.Age, User1.Address)
//
//	c.JSON(200, gin.H{
//		"code":    200,
//		"id":      User1.ID,
//		"name":    User1.Name,
//		"age":     User1.Age,
//		"address": User1.Address,
//	})
//	//db:= initDB()
//	//sqlDB, err := db.DB()
//	//
//	//// SetMaxIdleConns 设置空闲连接池中连接的最大数量
//	//sqlDB.SetMaxIdleConns(10)
//	//
//	//// SetMaxOpenConns 设置打开数据库连接的最大数量。
//	//sqlDB.SetMaxOpenConns(100)
//	//db.Create(&User1)
//	//_ = db.AutoMigrate()
//
//}
//
//type Users model {
//	ID      int    `json:"id" bind:"required" gorm:"id"` //必须添加属性字段
//	Name    string `json:"name" bind:"required" gorm:"name"`
//	Age     int    `json:"age" bind:"required" gorm:"age"`
//	Address string `json:"address" bind:"required" gorm:"address"`
//}
//
//func initDB() *gorm.DB {
//
//	dsn := "uep:1111111@tcp(172.16.50.52:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	if err != nil {
//		panic("数据库初始化失败 :" + err.Error())
//
//	}
//
//	return db
//}
//
////download	download modules to local cache(下载依赖包)
////edit	     edit go.mod from tools or scripts（编辑go.mod)
////graph	print module requirement graph (打印模块依赖图)
////init	initialize new module in current directory（在当前目录初始化mod）
////tidy	add missing and remove unused modules(拉取缺少的模块，移除不用的模块
////vendor	make vendored copy of dependencies(将依赖复制到vendor下)
////verify	verify dependencies have expected content (验证依赖是否正确）
////why	explain why packages or modules are needed(解释为什么需要依赖)
