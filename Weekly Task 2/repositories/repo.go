package repositories

import (
	"weekly-task-2/models"

	"github.com/labstack/echo"
)

type Database interface {
	GetItems() ([]models.Item, error)
	GetItemById(uuid string) (models.Item, error)
	CreateItem(item models.Item) error
	UpdateItem(echo.Context) error
	DeleteItem(uuid string) error
	GetByCategory(category_id uint) ([]models.Item, error)
	CreateCategory(category models.Category) error
	GetItemByName(item_name string) (models.Item, error)
}
