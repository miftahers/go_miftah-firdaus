package services

import (
	"weekly-task-3/dto"
	"weekly-task-3/models"

	"github.com/labstack/echo"
)

type BlogService interface {
	NewBlog(ctx echo.Context) error
	GetBlogs(ctx echo.Context) ([]models.Blog, error)
	GetBlogByID(ctx echo.Context) (models.Blog, error)
	UpdateBlog(ctx echo.Context) error
	DeleteBlog(ctx echo.Context) error
	GetBlogByCategory(ctx echo.Context) ([]models.Blog, error)
	GetBlogByTitle(ctx echo.Context) (models.Blog, error)
}

type UserService interface {
	SignUp(ctx echo.Context) error
	Login(ctx echo.Context) (dto.LoginUser, error)
}

type CategoryService interface {
	NewCategory(ctx echo.Context) error
}
