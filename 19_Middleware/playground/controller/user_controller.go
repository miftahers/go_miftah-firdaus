package controller

import (
	"net/http"
	"playground/config"
	m "playground/middleware"
	"playground/model"

	"github.com/labstack/echo"
)

func GetUserController(c echo.Context) error {
	var users []model.User

	// GORM akan mengecek tabel users dan menyimpan data-datanya ke variabel users
	err := config.DB.Find(&users).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    users,
	})
}

// controller
func CreateUserController(c echo.Context) error {
	// buat variable struct User
	user := model.User{}
	// bind data yang masuk ke struct User
	c.Bind(&user)
	//Insert data to DB
	err := config.DB.Save(&user).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create!",
		"data":    user,
	})
}

func LoginUserController(ctx echo.Context) error {
	// buat variable struct User
	user := model.User{}
	// bind data yang masuk ke struct User
	ctx.Bind(&user)
	//Cek data di model.User{}
	err := config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Fail Login",
			"error":   err.Error(),
		})
	}

	token, err := m.CreateToken(user.ID, user.Name)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Fail Login",
			"error":   err.Error(),
		})
	}

	userResponse := model.UserResponse{user.ID, user.Name, user.Email, token}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success login!",
		"data":    userResponse,
	})
}
