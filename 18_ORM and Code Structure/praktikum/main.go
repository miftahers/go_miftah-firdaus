package main

import (
	"praktikum/config"
	"praktikum/router"
)

func main() {
	config.InitDB()

	e := router.New()
	e.Start(":8080")
}
