package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialect/msyql"
)

var (
	db * gorm.DB
)

func Connect(){
	d, err := gorm.Open("mysql", "admin:Incorrect/go_bookstore?charset=utf8parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB{
	return db
}
