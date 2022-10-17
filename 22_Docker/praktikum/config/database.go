package config

import (
	"fmt"
	"praktikum/model"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	db, err := ConnectDB()
	if err != nil {
		time.Sleep(15 * time.Second)
		db, err = ConnectDB()
		if err != nil {
			panic(err)
		}
	}

	err = MigrateDB(db)
	if err != nil {
		panic(err)
	}
	return db
}

const (
	DB_Username string = "test_user"
	DB_Password string = "secret"
	DB_Port     string = "3306"
	DB_Host     string = "db-mysql"
	DB_Name     string = "test_database"
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
