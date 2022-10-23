package repositories

import (
	"weekly-task-3/models"

	"github.com/labstack/echo"
)

type Database interface {
	SignUp(user models.User) error
	Login(ctx echo.Context) (models.User, error)
	GetBlogs() ([]models.Blog, error)
	GetBlogByID(uuid string) (models.Blog, error)
	NewBlog(blog models.Blog) error
	NewCategory(category models.Category) error
	UpdateBlog(blog models.Blog) error
	DeleteBlog(uuid string) error
	GetBlogByCategory(category_id uint) ([]models.Blog, error)
	GetBlogByTitle(title string) (models.Blog, error)
}
