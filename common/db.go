package common

import "github.com/jinzhu/gorm"
import _ "github.com/go-sql-driver/mysql"
var(
	db *gorm.DB
)

//载入时自动执行
func init() {
	db,_ = gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/go_micro?charset=utf8&parseTime=True&loc=Local")

}

func GetDB()*gorm.DB  {
	return db
}