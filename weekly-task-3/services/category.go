package services

import (
	"weekly-task-3/models"

	"github.com/labstack/echo"
)

func (s *CategoryServ) NewCategory(ctx echo.Context) error {
	var category models.Category

	ctx.Bind(&category)

	err := s.Database.NewCategory(category)
	if err != nil {
		return err
	}
	return nil
}
