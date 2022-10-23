package services

import (
	"errors"
	"weekly-task-3/dto"
	"weekly-task-3/middleware"
	"weekly-task-3/models"

	"github.com/labstack/echo"
)

func (s *UserServ) SignUp(ctx echo.Context) error {
	var user models.User
	ctx.Bind(&user)
	if user.Username == "" {
		return errors.New("username should not empty")
	}
	err := s.Database.SignUp(user)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserServ) Login(ctx echo.Context) (dto.LoginUser, error) {
	user, err := s.Database.Login(ctx)
	if err != nil {
		return dto.LoginUser{}, err
	}

	var u dto.LoginUser
	u.ID = user.ID
	u.Username = user.Username

	token, err := middleware.GetToken(u.ID, u.Username)
	if err != nil {
		return dto.LoginUser{}, err
	}
	u.Token = token

	return u, nil
}
