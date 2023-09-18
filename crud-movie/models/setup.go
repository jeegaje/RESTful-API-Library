package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:3306)/movie_go?parseTime=true"))

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Movies{})

	DB = db
}
