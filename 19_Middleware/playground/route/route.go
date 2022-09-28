package route

import (
	"playground/constants"
	"playground/controller"
	m "playground/middleware"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	// routing
	e.GET("/users", controller.GetUserController)
	m.LogMiddleware(e)
	e.POST("/users", controller.CreateUserController)
	e.POST("/login", controller.LoginUserController)

	eAuthBasic := e.Group("/auth")
	eAuthBasic.Use(middleware.BasicAuth(m.BasicAuthDB))
	eAuthBasic.GET("/users", controller.GetUserController)

	eJwt := e.Group("/jwt")
	eJwt.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	eJwt.GET("/users", controller.GetUserController)

	return e
}
