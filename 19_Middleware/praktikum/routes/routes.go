package routes

import (
	"praktikum/constants"
	c "praktikum/controller"
	m "praktikum/middleware"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	// TODO handle TrailingSlash
	e.Pre(middleware.RemoveTrailingSlash())

	// TODO Implement logger
	m.LogMiddleware(e)

	//user routes
	users := e.Group("/users")
	users.POST("", c.CreateUserController)
	users.POST("/login", c.LoginUserController)
	//user need Auth
	users.GET("", c.GetUsersController, middleware.JWT([]byte(constants.SECRET_JWT_TOKEN)))
	users.GET("/:id", c.GetUserController, middleware.JWT([]byte(constants.SECRET_JWT_TOKEN)))
	users.PUT("/:id", c.UpdateUserController, middleware.JWT([]byte(constants.SECRET_JWT_TOKEN)))
	users.DELETE("/:id", c.DeleteUserController, middleware.JWT([]byte(constants.SECRET_JWT_TOKEN)))

	//book routes
	books := e.Group("/books")
	books.GET("", c.GetBooksController)
	books.GET("/:id", c.GetBookController)
	//book Need auth
	books.POST("", c.CreateBookController, middleware.JWT([]byte(constants.SECRET_JWT_TOKEN)))
	books.PUT("/:id", c.UpdateBookController, middleware.JWT([]byte(constants.SECRET_JWT_TOKEN)))
	books.DELETE("/:id", c.DeleteBookController, middleware.JWT([]byte(constants.SECRET_JWT_TOKEN)))

	return e
}
