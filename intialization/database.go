package intialization

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("mydb.db"), &gorm.Config{})
	if err != nil {
		panic("can't connect to database")
	}
}
