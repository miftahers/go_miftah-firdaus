package config

import (
	"fmt"

	"restful-api-testing/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

const (
	DB_Username string = "root"
	DB_Password string = ""
	DB_Port     string = "3306"
	DB_Host     string = "localhost"
	DB_Name     string = "crud_go"

	DB_Username_Test string = "root"
	DB_Password_Test string = ""
	DB_Port_Test     string = "3306"
	DB_Host_Test     string = "localhost"
	DB_Name_Test     string = "crud_go_test"
)

// REAL

func InitDB() {

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		DB_Username, DB_Password, DB_Host, DB_Port, DB_Name)

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	InitMigration()
}

func InitMigration() {
	DB.AutoMigrate(&model.User{}, &model.Book{})
}

// TEST

func InitDBTest() {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		DB_Username_Test, DB_Password_Test, DB_Host_Test, DB_Port_Test, DB_Name_Test)

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	InitMigrateTest()
}

func InitMigrateTest() {
	DB.Migrator().DropTable(&model.User{}, &model.Book{})
	DB.AutoMigrate(&model.User{}, &model.Book{})
}
