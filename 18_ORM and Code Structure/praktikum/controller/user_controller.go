package controller

import (
	"net/http"
	cfg "praktikum/config"
	"praktikum/model"
	"strconv"

	"github.com/labstack/echo"
)

func GetUsersController(ctx echo.Context) error {
	var users []model.User
	if err := cfg.DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

func GetUserController(ctx echo.Context) error {
	var user []model.User
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := cfg.DB.First(&user, "id = ?", id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user by id: " + strconv.Itoa(id),
		"user":    user,
	})
}

func CreateUserController(ctx echo.Context) error {
	user := model.User{}

	ctx.Bind(&user)

	if err := cfg.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusCreated, map[string]interface{}{
		"message": "User created!",
		"user":    user,
	})

}

func UpdateUserController(ctx echo.Context) error {
	var user model.User

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	name := ctx.FormValue("name")
	email := ctx.FormValue("email")
	password := ctx.FormValue("password")

	user.Name = name
	user.Email = email
	user.Password = password

	if err := cfg.DB.Save(&user).Where("id = ?", id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"messages": "user updated!",
		"user":     user,
	})
}

func DeleteUserController(ctx echo.Context) error {
	var user model.User

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := cfg.DB.Delete(&user, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"messages": "user with id: " + strconv.Itoa(id) + " deleted.",
	})
}
