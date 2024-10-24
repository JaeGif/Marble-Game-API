package handlers

import (
	"fmt"
	storage "marble-game-api/cmd/database"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetScores (c echo.Context) error {
	
	db := storage.GetDB()

	psqlStatement := `SELECT TOP 10 FROM scores`

	rows, err := db.Query(psqlStatement)

	if err != nil {
		// err exists
		return err
	}
	fmt.Println(rows)
	var scores = rows
	return c.JSON(http.StatusOK, scores)
}