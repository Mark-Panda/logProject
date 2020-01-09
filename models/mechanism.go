package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/smokezl/govalidators" //参数检验器
)

/**
机构表
 */
type Mechanism struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);not null;" validate:"required||string=5,30"`   //机构名字
	AccountNum string `gorm:"type:varchar(20);not null;" validate:"required||string=5,30"` //登录账号
	Password string `gorm:"type:varchar(20);not null;" validate:"required||string=5,30"` //登录密码
	Extend string `gorm:"type:varchar(20);not null;"`  //扩展字段
}


//创建新机构
func (mechan *Mechanism) createMechan(name, accountNum, pwd string) bool  {
	fmt.Println("机构参数", name)
	mechanism := &Mechanism{
		Name:   name,
		AccountNum: accountNum,
		Password: pwd,
		Extend: "",
	}
	validator := govalidators.New()
	errList := validator.Validate(mechanism)
	if errList != nil {
		for _, err := range errList {
			fmt.Println("参数验证err:", err)
			return false
		}
	}

	if err := GetDB().Create(mechanism).Error; err != nil {
		fmt.Println("createError:", err)
		return false
	}
	return true
}