package config

import (
	"fmt"
	"os"
	"praktikum/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	db, err := ConnectDB()
	if err != nil {
		panic(err)
	}

	err = MigrateDB(db)
	if err != nil {
		panic(err)
	}
	return db
}

var (
	DB_Username = os.Getenv("DB_USERNAME")
	DB_Password = os.Getenv("DB_PASSWORD")
	DB_Host     = os.Getenv("DB_HOST")
	DB_Port     = os.Getenv("DB_PORT")
	DB_Name     = os.Getenv("DB_NAME")
)

func ConnectDB() (*gorm.DB, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		DB_Username, DB_Password, DB_Host, DB_Port, DB_Name)

	return gorm.Open(mysql.Open(connectionString), &gorm.Config{})
}

func MigrateDB(db *gorm.DB) error {
	return db.AutoMigrate(
		model.User{},
	)
}
