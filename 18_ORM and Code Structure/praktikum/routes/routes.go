package routes

import (
	c "praktikum/controller"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	//router user
	users := e.Group("/users")
	users.GET("", c.GetUsersController)
	users.GET("/:id", c.GetUserController)
	users.POST("", c.CreateUserController)
	users.PUT("/:id", c.UpdateUserController)
	users.DELETE("/:id", c.DeleteUserController)

	//router book
	books := e.Group("/books")
	books.GET("", c.GetBooksController)
	books.GET("/:id", c.GetBookController)
	books.POST("", c.CreateBookController)
	books.PUT("/:id", c.UpdateBookController)
	books.DELETE("/:id", c.DeleteBookController)

	return e
}
