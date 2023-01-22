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
	// the username, password and database name used for gorm.Open() is dependent on the local environment
	// make your own database in your mysql server and edit the 3 variables here that are environment specific 
	d, err := gorm.Open("mysql", "thapelo:thapelo@/go-bookstore?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}