package main

import (
	storage "marble-game-api/cmd/database"
	"marble-game-api/cmd/handlers"

	//	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", handlers.Home)

	// connect to DB
	storage.InitDB()

	// starting routes
/* 	e.POST("/users", handlers.CreateUser)
	e.POST("/measurements", handlers.CreateMeasurement)
	e.PUT("/users/:id", handlers.UpdateUser)
	e.PUT("/measurements/:id", handlers.UpdateMeasurement) */

	e.GET("/scores", handlers.GetScores)
	e.Use(handlers.LogRequest)
//	e.Use(middleware.Logger)
/* e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	AllowOrigins: []string{"http://localhost:3000"},
	AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
  })) */
	// start server or log fatal
	e.Logger.Fatal(e.Start(":8080"))
	
}