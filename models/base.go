package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB //database

func init() {
	conn, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	db = conn
	//Migarate db
	db.Debug().AutoMigrate(&User{})
}
func GetDB() *gorm.DB {
	return db
}
