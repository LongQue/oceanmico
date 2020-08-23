package main

import "github.com/jinzhu/gorm"

func main() {
	_, err := gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/go_micro?charset=utf8&parseTime=True&loc=Local")
	if err!= nil {
		panic(err)
	}
}
