package controller

import (
	"net/http"
	cfg "praktikum/config"
	m "praktikum/middleware"
	"praktikum/model"

	"github.com/labstack/echo"
)

func LoginUserController(ctx echo.Context) error {

	user := model.User{}
	ctx.Bind(&user)

	err := cfg.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Fail Login",
			"error":   err.Error(),
		})
	}

	// TODO Create JWT Token
	token, err := m.CreateToken(int(user.ID), user.Name)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Fail Login",
			"error":   err.Error(),
		})
	}

	userResponse := model.UserResponse{int(user.ID), user.Name, user.Email, token}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success login!",
		"data":    userResponse,
	})
}
