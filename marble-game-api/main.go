package main

import (
	storage "marble-game-api/cmd/database"
	"marble-game-api/cmd/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Debug = true
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	storage.InitDB()

	e.GET("/", handlers.Home)

	// connect to DB

	// starting routes
	/* 	e.POST("/users", handlers.CreateUser)
	   	e.POST("/measurements", handlers.CreateMeasurement)
	   	e.PUT("/users/:id", handlers.UpdateUser)
	   	e.PUT("/measurements/:id", handlers.UpdateMeasurement) */
	e.POST("/scores", handlers.CreateScore)

	e.GET("/scores", handlers.GetScores)
	//	e.Use(handlers.LogRequest)
	//	e.Use(middleware.Logger)
	/* e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	  })) */
	// start server or log fatal
	e.Logger.Fatal(e.Start(":8080"))

}
