package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	dbConn *gorm.DB
)

func Connect() {
	db, err := gorm.Open("mysql", "root:admin@tcp(localhost:3306)/bookstore?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	dbConn = db
}

func GetDBConn() *gorm.DB {
	return dbConn
}
