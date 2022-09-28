package middleware

import (
	"playground/config"
	"playground/model"

	"github.com/labstack/echo"
)

func BasicAuthDB(username, password string, c echo.Context) (bool, error) {
	// Be careful to use constant time comparison to prevent timing attacks
	var user model.User
	err := config.DB.Where("email= ? AND password= ?", username, password).First(&user).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
