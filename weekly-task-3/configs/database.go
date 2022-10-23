package configs

import (
	"fmt"
	"weekly-task-3/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB_Username string = "root"
	DB_Password string = ""
	DB_Port     string = "3306"
	DB_Host     string = "localhost"
	DB_Name     string = "weekly_task_3"
)

func ConnectDB() (*gorm.DB, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		DB_Username, DB_Password, DB_Host, DB_Port, DB_Name)

	return gorm.Open(mysql.Open(connectionString), &gorm.Config{})
}

func MigrateDB(db *gorm.DB) error {
	return db.AutoMigrate(
		models.Category{},
		models.User{},
		models.Blog{},
	)
}
