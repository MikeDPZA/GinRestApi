package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var ctx *gorm.DB

func Init() {
	dsn := "root:password@tcp(127.0.0.1:3306)/GinBookStore?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database" + err.Error())
	}
	ctx = db
}

func Context() *gorm.DB {
	return ctx
}
