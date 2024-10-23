package main

import (
	storage "marble-game-api/cmd/database"
	"marble-game-api/cmd/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", handlers.Home)

	// connect to DB
	storage.InitDB()

	// start server or log fatal
	e.Logger.Fatal(e.Start(":8080"))
	
}