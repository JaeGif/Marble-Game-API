package handlers

import (
	"fmt"
	storage "marble-game-api/cmd/database"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetScores (c echo.Context) error {
	fmt.Println("Entry")
	
	db := storage.GetDB()
	if db == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Database connection failed")
	}
	
	psqlStatement := `SELECT * FROM scores LIMIT 10`

	rows, err := db.Query(psqlStatement)
	fmt.Println(err, rows)
	if err != nil {
		// err exists
		return err
	}

	fmt.Println(rows)
	var scores = rows
	fmt.Println(http.StatusOK, scores)
	return c.JSON(http.StatusOK, scores) 
}