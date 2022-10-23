package controller

import (
	"net/http"
	"weekly-task-3/services"

	"github.com/labstack/echo"
)

type UserHandler struct {
	services.UserService
}

func (h *UserHandler) SignUp(ctx echo.Context) error {
	err := h.UserService.SignUp(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, map[string]interface{}{
		"message": "created",
	})
}

func (h *UserHandler) Login(ctx echo.Context) error {
	result, err := h.UserService.Login(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    result,
	})
}
