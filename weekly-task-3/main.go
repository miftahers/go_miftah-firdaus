package main

import (
	"weekly-task-3/configs"
	"weekly-task-3/routes"
)

func main() {
	db, err := configs.ConnectDB()
	if err != nil {
		panic(err)
	}

	err = configs.MigrateDB(db)
	if err != nil {
		panic(err)
	}

	e := routes.New(db)
	e.Start(":8080")
}
