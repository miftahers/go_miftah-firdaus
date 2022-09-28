package config

import (
	"playground/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:@tcp(localhost:3306)/playground?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	// init migrate di init DB untuk membuat table baru
	InitMigrate()
}

func InitMigrate() {
	//Dapat membuat structur table User secara otomatis jika table user belum ada didalam Database
	DB.AutoMigrate(&model.User{})
}
