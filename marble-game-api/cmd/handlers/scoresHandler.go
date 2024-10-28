package handlers

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func GetScores (c echo.Context) error {
	fmt.Println("Entry")

	return c.String(200, "Done")

/* 	db := storage.GetDB()

	psqlStatement := `SELECT TOP 10 FROM scores`

	rows, err := db.Query(psqlStatement)
	fmt.Println(err, rows)
	if err != nil {
		// err exists
		return err
	}

	fmt.Println(rows)
	var scores = rows
	fmt.Println(http.StatusOK, scores)
	return c.JSON(http.StatusOK, scores) */
}