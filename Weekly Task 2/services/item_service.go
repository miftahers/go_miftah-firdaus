package services

import (
	"errors"
	"strconv"
	"weekly-task-2/models"
	"weekly-task-2/repositories"

	"github.com/google/uuid"
	"github.com/labstack/echo"
)

type ItemService interface {
	GetItems() ([]models.Item, error)
	GetItemById(ctx echo.Context) (models.Item, error)
	CreateItem(ctx echo.Context) error
	UpdateItem(ctx echo.Context) error
	DeleteItem(ctx echo.Context) error
	GetByCategory(ctx echo.Context) ([]models.Item, error)
	CreateCategory(ctx echo.Context) error
	GetItemByName(ctx echo.Context) (models.Item, error)
}

type itemServ struct {
	repositories.Database
}

func NewItemServices(db repositories.Database) ItemService {
	return &itemServ{
		Database: db,
	}
}

func (itemServ *itemServ) GetItems() ([]models.Item, error) {
	result, err := itemServ.Database.GetItems()
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (itemServ *itemServ) GetItemById(ctx echo.Context) (models.Item, error) {

	uuid := ctx.Param("id")

	result, err := itemServ.Database.GetItemById(uuid)
	if err != nil {
		return models.Item{}, err
	}
	return result, nil
}
func (itemServ *itemServ) CreateItem(ctx echo.Context) error {
	var item models.Item

	ctx.Bind(&item)

	item.UUID = uuid.NewString()
	if item.Name == "" {
		return errors.New("field name cannot be empty")
	}

	err := itemServ.Database.CreateItem(item)
	if err != nil {
		return err
	}
	return nil
}
func (itemServ *itemServ) UpdateItem(ctx echo.Context) error {
	err := itemServ.Database.UpdateItem(ctx)
	if err != nil {
		return err
	}
	return nil
}
func (itemServ *itemServ) DeleteItem(ctx echo.Context) error {

	uuid := ctx.Param("id")

	err := itemServ.Database.DeleteItem(uuid)
	if err != nil {
		return err
	}
	return nil
}
func (itemServ *itemServ) GetByCategory(ctx echo.Context) ([]models.Item, error) {

	category_id, err := strconv.Atoi(ctx.Param("category_id"))
	if err != nil {
		return nil, err
	}

	result, err := itemServ.Database.GetByCategory(uint(category_id))
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (itemServ *itemServ) CreateCategory(ctx echo.Context) error {
	var category models.Category

	ctx.Bind(&category)

	err := itemServ.Database.CreateCategory(category)
	if err != nil {
		return err
	}
	return nil
}
func (itemServ *itemServ) GetItemByName(ctx echo.Context) (models.Item, error) {

	item_name := ctx.QueryParam("keyword")

	result, err := itemServ.Database.GetItemByName(item_name)
	if err != nil {
		return models.Item{}, err
	}
	return result, nil
}
