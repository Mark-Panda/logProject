package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/smokezl/govalidators" //参数检验器
)

type User struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);not null;" validate:"required||string=5,30"`
	Phone string `gorm:"type:varchar(20);not null;" validate:"required||string=11"`
	Password string `gorm:"type:varchar(20);not null;" validate:"required||string=8,20"`
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
	user := &User{
		Name:name,
		Phone:phone,
		Password:pwd,
	}
	validator := govalidators.New()
	errList := validator.Validate(user)
	if errList != nil {
		for _, err := range errList {
			fmt.Println("参数验证err:", err)
			return false
		}
	}
	if err := GetDB().Create(user).Error; err != nil {
		fmt.Println("createError:", err)
		return false
	}
	return true
}

