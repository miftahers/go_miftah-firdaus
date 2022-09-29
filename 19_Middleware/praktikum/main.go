package main

import (
	"praktikum/config"
	"praktikum/routes"
)

func main() {
	//init Database
	config.InitDB()

	//routes
	e := routes.New()
	e.Start(":8080")
}
