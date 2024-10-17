package controllers

import (
	"net/http"
	"your-app/models"

	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
    users, err := models.GetAllUsers()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }
    return c.JSON(http.StatusOK, users)
}

func CreateUser(c echo.Context) error {
    var user models.User
    if err := c.Bind(&user); err != nil {
        return c.JSON(http.StatusBadRequest, err.Error())
    }

    if err := models.CreateUser(&user); err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }
    return c.JSON(http.StatusCreated, user)
}
