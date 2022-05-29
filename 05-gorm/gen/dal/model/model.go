package model

// User 定义User结构体
type User struct {
	UUID    string `gorm:"primaryKey;uuid"`
	Name    string `gorm:"name"`
	Age     int    `gorm:"age"`
	Version int    `gorm:"version"`
}
