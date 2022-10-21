package controller

import (
	"net/http"
	"weekly-task-2/services"

	"github.com/labstack/echo"
)

type ItemHandler struct {
	services.ItemService
}

func (h *ItemHandler) GetItems(ctx echo.Context) error {
	if ctx.QueryParam("keyword") == "" {
		result, err := h.ItemService.GetItems()
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"message": "success",
			"data":    result,
		})
	} else {
		result, err := h.ItemService.GetItemByName(ctx)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"message": "success",
			"data":    result,
		})
	}
}

func (h *ItemHandler) GetItemById(ctx echo.Context) error {
	item, err := h.ItemService.GetItemById(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    item,
	})
}

func (h *ItemHandler) CreateItem(ctx echo.Context) error {
	err := h.ItemService.CreateItem(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Created!",
	})
}

func (h *ItemHandler) UpdateItem(ctx echo.Context) error {
	err := h.ItemService.UpdateItem(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Updated!",
	})
}

func (h *ItemHandler) DeleteItem(ctx echo.Context) error {
	err := h.ItemService.DeleteItem(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "Deleted!",
	})
}

func (h *ItemHandler) GetByCategory(ctx echo.Context) error {
	items, err := h.ItemService.GetByCategory(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    items,
	})
}
func (h *ItemHandler) CreateCategory(ctx echo.Context) error {
	err := h.ItemService.CreateCategory(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Created!",
	})
}

// func (h *ItemHandler) GetItemByName(ctx echo.Context) error {
// 	item, err := h.ItemService.GetItemByName(ctx)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
// 	}
// 	return ctx.JSON(http.StatusOK, map[string]interface{}{
// 		"message": "Success",
// 		"data":    item,
// 	})
// }
