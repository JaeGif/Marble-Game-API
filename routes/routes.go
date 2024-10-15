package routes

import (
	"/controllers"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
    e.GET("/users", controllers.GetUsers)
    e.POST("/users", controllers.CreateUser)
    e.GET("/users/:id", controllers.GetUser)
    e.PUT("/users/:id", controllers.UpdateUser)
    e.DELETE("/users/:id", controllers.DeleteUser)
}
