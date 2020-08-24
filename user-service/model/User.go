package model

import "github.com/jinzhu/gorm"

type User struct {
	//gorm.Model存在常用的默认字段，如id 创建时间等
	gorm.Model
	Username string `gorm:"type:varchar(255);unique_index"`
	Password string `gorm:"type:varchar(255)"`
	Email    string `gorm:"type:varchar(100)"`
	Channel  int64  `gorm:"type:tinyint(2)"`
}
