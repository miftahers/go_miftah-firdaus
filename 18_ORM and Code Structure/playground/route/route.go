package route

import (
	"playground/controller"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	// routing
	e.GET("/users", controller.GetUserController)
	e.POST("/users", controller.CreateUserController)
	return e
}
