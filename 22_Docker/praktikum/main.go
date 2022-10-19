package main

import (
	"praktikum/config"
	"praktikum/routes"
)

func main() {

	config.InitConfig()

	db := config.InitDB()

	e := routes.Init(db)

	e.Start(config.Cfg.APIPort)
}
