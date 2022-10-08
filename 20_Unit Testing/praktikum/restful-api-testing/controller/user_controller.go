package controller

import (
	"net/http"

	cfg "restful-api-testing/config"
	m "restful-api-testing/middleware"
	"restful-api-testing/model"
	"strconv"

	"github.com/labstack/echo"
)

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
	// TODO Bind created token into user token
	user.Token = token

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Login user",
		"data":    user,
	})
}

func GetUsersController(ctx echo.Context) error {
	var users []model.User
	if err := cfg.DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// get Id from JWT
	// myUserId := middleware.ExtractTokenUserId(c)

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    users,
	})
}

func GetUserController(ctx echo.Context) error {
	var user model.User

	id, _ := strconv.Atoi(ctx.Param("id"))

	if err := cfg.DB.First(&user, "id = ?", id).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    user,
	})
}

func CreateUserController(ctx echo.Context) error {
	var user model.User

	ctx.Bind(&user)

	if err := cfg.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, map[string]interface{}{
		"message": "user created!",
		"data":    user,
	})
}

func UpdateUserController(ctx echo.Context) error {
	var user model.User

	id, _ := strconv.Atoi(ctx.Param("id"))

	name := ctx.FormValue("name")
	email := ctx.FormValue("email")
	password := ctx.FormValue("password")

	if err := cfg.DB.Model(&user).Where("users.id = ?", id).Updates(model.User{Name: name, Email: email, Password: password}).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "user updated!",
	})
}

func DeleteUserController(ctx echo.Context) error {
	var user model.User

	id, _ := strconv.Atoi(ctx.Param("id"))

	if err := cfg.DB.Delete(&user, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	/*
		untuk recover data abis soft delete tinggal set column deleted_at ke NULL

		to delete permanently
		if err := cfg.DB.Unscoped().Delete(&user, id).Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
	*/
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "delete success",
	})
}
