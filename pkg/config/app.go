package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Look up what this block does
var (
	db * gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "thapelo:thapelo@/go-bookstore?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}