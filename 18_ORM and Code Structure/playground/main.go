package main

import (
	"playground/config"
	"playground/route"
)

func main() {
	config.InitDB()

	e := route.New()
	e.Start(":8080")
}
