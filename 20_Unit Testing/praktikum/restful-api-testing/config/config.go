package config

import (
	"fmt"
	"os"
	"restful-api-testing/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	DB_Username string = "root"
	DB_Password string = ""
	DB_Port     string = "3306"
	DB_Host     string = "localhost"
	DB_Name     string = "crud_go"
)

var (
	Cfg *Config
	DB  *gorm.DB
)

type Config struct {
	APIPort     string
	APIKey      string
	TokenSecret string
}

func Init() {
	InitConfig()
	InitDB()
}

func InitConfig() {
	cfg := &Config{}

	cfg.APIPort = os.Getenv("APIPort")
	cfg.APIKey = os.Getenv("APIKey")
	cfg.TokenSecret = "SuperIdol"

	Cfg = cfg
}

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
