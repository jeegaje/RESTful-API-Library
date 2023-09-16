package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func conn() {

	fmt.Println("hellp")
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:3306)/phabricator_user"))

	if err != nil {
		panic(err)
	}

	DB = db
}
