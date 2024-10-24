package handlers

import (
	"marble-game-api/cmd/models"
	"marble-game-api/cmd/repositories"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateMeasurement(c echo.Context) error {
  measurement := models.Measurements{}
  c.Bind(&measurement)
  newMeasurement, err := repositories.CreateMeasurement(measurement)
  if err != nil {
    return c.JSON(http.StatusInternalServerError, err.Error())
  }
  return c.JSON(http.StatusCreated, newMeasurement)
}
