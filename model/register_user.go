package model

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

////https://www.cnblogs.com/jiujuan/p/13823864.html

type Users struct {
	ID       uint   `json:"id"  gorm:"primarykey;not null;autoIncrement" ` //必须添加属性字段
	UUID     string `json:"uuid"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Age      uint   `json:"age"`
	Address  string `json:"address"`
	Token    bool   `json:"token"`
	gorm.Model
}


//SelectWritingUser 通过查询姓名写入到结构体中
func (user *Users) SelectUserExiting(db *gorm.DB) bool {
	//db.Where("name = ?", user.Name).First(&user)

	if result := db.Debug().Raw("select name, password from users where name = ? and deleted_at is null", user.Name).Scan(user); result.Error == nil {

		//db.Debug().Model(&Users{}).Find(&SelectUsers{}).Where("name = ?", user.Name)
		//db.Debug().Select("name", "password").Model(&Users{}).Where("name = ?", user.Name).Scan(user)
		//db.Debug().Raw("select name, password from users where name = ? and deleted_at is null;
		//log.Printf("用户存在，共找到 %v 条数据 ", result.RowsAffected)

		return true
	}
	return false
}

func (user *Users) SetUserUuid() {
	var err error
	u1 := uuid.Must(uuid.NewV4(), err).String()
	fmt.Printf("UUIDv4: %s\n", u1)
	user.UUID = u1
}

