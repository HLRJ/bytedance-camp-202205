package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

// User 定义User结构体
type User struct {
	UUID    string `gorm:"uuid"`
	Name    string `gorm:"name"`
	Age     int    `gorm:"age"`
	Version int    `gorm:"version"`
}

//https://gorm.io/zh_CN/docs/

func main() {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// 迁移 schema
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Println(err)
		return
	}

}
