package models

import "github.com/jinzhu/gorm"

type Mechanism struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);not null;"`
	Extend string `gorm:"type:varchar(20);not null;"`
}