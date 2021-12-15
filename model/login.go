package model

//修改了登录结构体只需要传输name，password。
//遗留问题，传输其他数据也能绑定，需要增加限制
//密码加密失败问题

type LoginUsers struct {
	Name     string `json:"name"  binding:"required,gt=3,lt=6" gorm:"name;primary_key;not null" ` // 10 <= age <= 120。
	Password string `json:"password"  binding:"required,gt=4,lt=10"`
}
