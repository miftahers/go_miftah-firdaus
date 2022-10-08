package main

import (
	"restful-api-testing/config"
	"restful-api-testing/routes"
)

func main() {
	//init Database
	config.Init()

	//routes
	e := routes.New()
	e.Start(config.Cfg.APIPort)
}
