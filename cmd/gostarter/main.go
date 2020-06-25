package main

import (
	"GoStarter/internal/add"
	"GoStarter/pkg/app"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	appConfig := app.Config{
		DBType:       "sqlite3",
		DBConnString: "./database.sqlite",
		Host:         ":5000",
	}

	app, err := app.CreateApp(appConfig)
	if err != nil {
		log.Fatal(err)
	}

	app.GET("/api/v1/add", add.Handler)

	app.Run()
}
