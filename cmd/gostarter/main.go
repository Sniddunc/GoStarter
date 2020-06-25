package main

import (
	"GoStarter/internal/add"
	"GoStarter/pkg/app"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	appConfig := app.Config{
		DBType:       "sqlite3",
		DBConnString: "./database.sqlite",
		Host:         ":5000",
		DevMode:      os.Getenv("MODE") == "development",
	}

	app, err := app.CreateApp(appConfig)
	if err != nil {
		log.Fatal(err)
	}

	// Route Declarations
	app.GET("/api/v1/add", add.Handler)

	app.Run()
}
