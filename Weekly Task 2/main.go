package main

import (
	"weekly-task-2/configs"
	"weekly-task-2/routes"
)

func main() {

	db := configs.InitDB()
	e := routes.New(db)
	e.Start(":8080")
}
