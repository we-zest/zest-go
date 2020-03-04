package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"zest-go/config"
)

var DB *gorm.DB

func init() {
	mysql := config.Mysql
	args := mysql["username"] + ":" + mysql["password"] + "@tcp(" + mysql["host"] + ":" + mysql["port"] + ")/" + mysql["database"]
	db, err := gorm.Open("mysql", args+"?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Panicln("mysql err:", err.Error())
	}

	DB = db
}
