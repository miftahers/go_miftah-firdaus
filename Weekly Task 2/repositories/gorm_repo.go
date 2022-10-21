package repositories

import (
	"net/http"
	"strconv"
	"weekly-task-2/models"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type GormSQL struct {
	DB *gorm.DB
}

func NewGorm(db *gorm.DB) Database {
	return &GormSQL{
		DB: db,
	}
}

func (gdb *GormSQL) GetItems() ([]models.Item, error) {
	var items []models.Item
	if err := gdb.DB.Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (gdb *GormSQL) GetItemById(uuid string) (models.Item, error) {
	var item models.Item

	if err := gdb.DB.First(&item, "uuid = ?", uuid).Error; err != nil {
		return models.Item{}, err
	}
	return item, nil
}

func (gdb *GormSQL) CreateItem(item models.Item) error {

	if err := gdb.DB.Save(&item).Error; err != nil {
		return err
	}
	return nil
}

func (gdb *GormSQL) UpdateItem(ctx echo.Context) error {

	var (
		oldItem, newItem                                     models.Item
		name, description, stockStr, priceStr, categoryIDStr string
		stock, price                                         int
		categoryID                                           uint
	)

	uuid := ctx.Param("id")

	if err := gdb.DB.First(&oldItem, "uuid = ?", uuid).Error; err != nil {
		return err
	}

	name = ctx.FormValue("name")
	description = ctx.FormValue("description")
	stockStr = ctx.FormValue("stock")
	priceStr = ctx.FormValue("price")
	categoryIDStr = ctx.FormValue("category_id")

	if name == "" {
		name = oldItem.Name
	}
	if description == "" {
		description = oldItem.Description
	}
	if categoryIDStr == "" {
		categoryID = oldItem.CategoryID
	} else {
		v, err := strconv.Atoi(categoryIDStr)
		if err != nil {
			return err
		}
		categoryID = uint(v)
	}
	if stockStr == "" {
		stock = oldItem.Stock
	} else {
		var err error
		stock, err = strconv.Atoi(stockStr)
		if err != nil {
			return err
		}
	}
	if priceStr == "" {
		price = oldItem.Price
	}

	if err := gdb.DB.Model(&newItem).Where("items.uuid = ?", uuid).Updates(
		models.Item{
			Name:        name,
			Description: description,
			Stock:       stock,
			Price:       price,
			CategoryID:  categoryID,
		}).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}

func (gdb *GormSQL) DeleteItem(uuid string) error {
	var item models.Item

	if err := gdb.DB.Unscoped().Where("uuid = ?", uuid).Delete(&item).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return nil
}

func (gdb *GormSQL) GetByCategory(category_id uint) ([]models.Item, error) {
	var items []models.Item

	if err := gdb.DB.Model(&models.Item{}).Where("category_id = ?", category_id).Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (gdb *GormSQL) CreateCategory(category models.Category) error {
	if err := gdb.DB.Save(&category).Error; err != nil {
		return err
	}
	return nil
}

func (gdb *GormSQL) GetItemByName(item_name string) (models.Item, error) {
	var item models.Item

	if err := gdb.DB.Where("name = ?", item_name).First(&item).Error; err != nil {
		return models.Item{}, err
	}
	return item, nil
}
