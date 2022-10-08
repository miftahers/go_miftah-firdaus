package routes

import (
	"restful-api-testing/config"
	"restful-api-testing/constants"
	"restful-api-testing/controller"
	mid "restful-api-testing/middleware"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	// TODO handle TrailingSlash
	e.Pre(middleware.RemoveTrailingSlash())

	// TODO Implement logger
	mid.LogMiddleware(e)

	//user routes
	users := e.Group(constants.UserIndex)
	users.POST("", controller.CreateUserController)
	users.POST("/login", controller.LoginUserController)
	users.GET("", controller.GetUsersController, middleware.JWT([]byte(config.Cfg.TokenSecret)))    // Tested
	users.GET("/:id", controller.GetUserController, middleware.JWT([]byte(config.Cfg.TokenSecret))) // Tested
	users.PUT("/:id", controller.UpdateUserController, middleware.JWT([]byte(config.Cfg.TokenSecret)))
	users.DELETE("/:id", controller.DeleteUserController, middleware.JWT([]byte(config.Cfg.TokenSecret)))

	//book routes
	books := e.Group(constants.BookIndex)
	books.GET("", controller.GetBooksController)
	books.GET("/:id", controller.GetBookController)
	books.POST("", controller.CreateBookController, middleware.JWT([]byte(config.Cfg.TokenSecret)))
	books.PUT("/:id", controller.UpdateBookController, middleware.JWT([]byte(config.Cfg.TokenSecret)))
	books.DELETE("/:id", controller.DeleteBookController, middleware.JWT([]byte(config.Cfg.TokenSecret)))

	return e
}
