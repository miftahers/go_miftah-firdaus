package routes

import (
	controller "weekly-task-2/controllers"
	"weekly-task-2/repositories"
	"weekly-task-2/services"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gorm.io/gorm"
)

func New(db *gorm.DB) *echo.Echo {
	e := echo.New()
	// TODO handle TrailingSlash
	e.Pre(middleware.RemoveTrailingSlash())

	// TODO Implement logger
	middleware.Logger()

	repo := repositories.NewGorm(db)
	service := services.NewItemServices(repo)
	handler := controller.ItemHandler{
		ItemService: service,
	}
	items := e.Group("/items")
	items.GET("", handler.GetItems)
	items.GET("/:id", handler.GetItemById)
	items.POST("", handler.CreateItem)
	items.PUT("/:id", handler.UpdateItem)
	items.DELETE("/:id", handler.DeleteItem)
	category := items.Group("/category")
	category.POST("", handler.CreateCategory)
	category.GET("/:category_id", handler.GetByCategory)
	// items.GET("", handler.GetItemByName)
	return e
}
