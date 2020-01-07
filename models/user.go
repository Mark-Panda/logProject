package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	//ID   string `gorm:"type:int(11);not null;"`// 名为`ID`的字段会默认作为表的主键
	Name string `gorm:"type:varchar(20);not null;"`
	Phone string `gorm:"type:varchar(20);not null;"`
	Password string `gorm:"type:varchar(20);not null;"`
}

func (use *User) FindOneByOps(phone string) *User {
	fmt.Println("参数", phone)
	user := User{}
	if err := GetDB().Where("phone = ?", phone).Find(&user).Error; err != nil {
		fmt.Println("error:", err)
		return nil
	}
	fmt.Println("数据库", user)
	return &user
}

func (use *User) CheckUserInfo(name, pwd string) bool  {
	fmt.Println("验证参数", name, pwd)
	user := User{}

	if err := GetDB().Where("name = ?", name).Find(&user).Error; err != nil {
		fmt.Println("查询失败", err)
		return false
	}
	fmt.Println("患者信息", user)
	if user.Password == pwd {
		return true
	}else {
		return false
	}

}

func (use *User) CreateUser(phone, name, pwd string) bool {
	fmt.Println("参数1111", phone, name)
	user := User{
		Name:name,
		Phone:phone,
		Password:pwd,
	}
	if err := GetDB().Create(&user).Error; err != nil {
		fmt.Println("createError:", err)
		return false
	}
	return true
}



// 禁用默认表名的复数形式，如果置为 true，则 `User` 的默认表名是 `user`
//db.SingularTable(true)