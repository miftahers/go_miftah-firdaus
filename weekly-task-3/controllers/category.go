package controller

import (
	"net/http"
	"weekly-task-3/services"

	"github.com/labstack/echo"
)

type CategoryHandler struct {
	services.CategoryService
}

func (h *CategoryHandler) NewCategory(ctx echo.Context) error {
	err := h.CategoryService.NewCategory(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusCreated, map[string]interface{}{
		"message": "created",
	})
}
