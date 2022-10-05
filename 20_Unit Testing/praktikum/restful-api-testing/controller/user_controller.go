package controller

import (
	"net/http"

	cfg "restful-api-testing/config"
	m "restful-api-testing/middleware"
	"restful-api-testing/model"
	"strconv"

	"github.com/labstack/echo"
)

type UserResponse struct {
	Message string
	Data    []model.User
}

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

	if err := cfg.DB.Model(&user).Where("users.id = ?", id).Updates(model.User{Name: name, Email: email, Password: password}).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"messages": "user updated!",
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
	/*
		untuk recover data abis soft delete tinggal set column deleted_at ke NULL

		to delete permanently
		if err := cfg.DB.Unscoped().Delete(&user, id).Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
	*/
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"messages": "user with id: " + strconv.Itoa(id) + " deleted.",
	})
}

func LoginUserController(ctx echo.Context) error {
	user := model.User{}
	ctx.Bind(&user)

	err := cfg.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Fail Login user",
			"error":   err.Error(),
		})
	}

	// TODO Create JWT Token
	token, err := m.CreateToken(int(user.ID), user.Name)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Fail create JWT Token",
			"error":   err.Error(),
		})
	}

	user.Token = token

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success login",
		"user":    user,
	})
}
