package main

import (
	"restful-api-testing/config"
	"restful-api-testing/routes"
)

func main() {
	//init Database
	config.InitDB()

	//routes
	e := routes.New()
	e.Start(":8080")
}
