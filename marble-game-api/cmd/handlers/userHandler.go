package handlers

import (
	"marble-game-api/cmd/models"
	"marble-game-api/cmd/repositories"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateUser(c echo.Context) error {
  user := models.User{}
  c.Bind(&user)
  newUser, err := repositories.CreateUser(user)
  if err != nil {
    return c.JSON(http.StatusInternalServerError, err.Error())
  }
  return c.JSON(http.StatusCreated, newUser)
}
