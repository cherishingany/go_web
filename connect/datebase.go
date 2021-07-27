package connect

import (
	"fmt"
	"github.com/spf13/viper"
	"go_web/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

func InitDB() *gorm.DB {

	ConfigInit()

	//dsn := "uep:1111111@tcp(172.16.50.52:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&loc=Local",
		viper.Get("database.user"),
		viper.Get("database.password"),
		viper.Get("database.host"),
		viper.Get("database.port"),
		viper.Get("database.database"),
		viper.Get("database.charset"),
	)
	//log.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic("Gorm 异常：" + err.Error())
	}
	//绑定数据库模型
	if err = db.AutoMigrate(&model.Users{});err != nil {   //db.AutoMigrate(&model.Users{}，&model.Videos{})
		panic("AutoMigrate 异常："+ err.Error())
	}

	//根据*grom.DB对象获得*sql.DB的通用数据库接口
	sqlDb, _ := db.DB()
	sqlDb.SetMaxIdleConns(10) //设置最大连接数
	sqlDb.SetMaxOpenConns(10) //设置最大的空闲连接数
	return db
}

func ConfigInit() {
	workspace, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workspace + "/config")

	if err := viper.ReadInConfig(); err != nil {
		panic(workspace + "读取数据库配置文件出错-----" + err.Error())
	}

}
