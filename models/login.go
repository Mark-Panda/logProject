package models

import "github.com/jinzhu/gorm"

type loginError struct {
	msg string
}

type Login struct {
	gorm.Model
	Phone string `gorm:"phone"`
	Password string `json:"password"`
	Msg loginError `json:"msg"`
}


func (l *Login) readError() *Login {

	return l
}

func (l *Login) writeError(c *Login) bool  {
	return true
}
