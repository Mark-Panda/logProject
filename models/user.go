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

func (use *User) CreateUser(phone, name string) bool {
	fmt.Println("参数1111", phone, name)
	user := User{
		Name:name,
		Phone:phone,
	}
	if err := GetDB().Create(&user).Error; err != nil {
		fmt.Println("createError:", err)
		return false
	}
	return true
}



// 禁用默认表名的复数形式，如果置为 true，则 `User` 的默认表名是 `user`
//db.SingularTable(true)