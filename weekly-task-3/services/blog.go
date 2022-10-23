package services

import (
	"strconv"
	"weekly-task-3/models"

	"github.com/google/uuid"
	"github.com/labstack/echo"
)

func (s *BlogServ) GetBlogs(ctx echo.Context) ([]models.Blog, error) {
	blogs, err := s.Database.GetBlogs()
	if err != nil {
		return nil, err
	}

	return blogs, nil
}

func (s *BlogServ) GetBlogByID(ctx echo.Context) (models.Blog, error) {
	uuid := ctx.Param("id")
	blog, err := s.Database.GetBlogByID(uuid)
	if err != nil {
		return models.Blog{}, err
	}
	return blog, nil
}

func (s *BlogServ) NewBlog(ctx echo.Context) error {
	var blog models.Blog

	// /* tadinya buat ngambil user_id
	t, err := DecodeJWT(ctx)
	if err != nil {
		return err
	}
	// */

	ctx.Bind(&blog)

	blog.UserID = t.UserID
	blog.UUID = uuid.NewString()

	err = s.Database.NewBlog(blog)
	if err != nil {
		return err
	}
	return nil
}

func (s *BlogServ) UpdateBlog(ctx echo.Context) error {

	var new models.Blog

	ctx.Bind(&new)

	new.UUID = ctx.Param("id")

	err := s.Database.UpdateBlog(new)
	if err != nil {
		return err
	}
	return nil
}

func (s *BlogServ) DeleteBlog(ctx echo.Context) error {
	uuid := ctx.Param("id")
	err := s.Database.DeleteBlog(uuid)
	if err != nil {
		return err
	}
	return nil
}

func (s *BlogServ) GetBlogByCategory(ctx echo.Context) ([]models.Blog, error) {
	category_id, err := strconv.Atoi(ctx.Param("category_id"))
	if err != nil {
		return nil, err
	}
	blogs, err := s.Database.GetBlogByCategory(uint(category_id))
	if err != nil {
		return nil, err
	}

	return blogs, nil
}

func (s *BlogServ) GetBlogByTitle(ctx echo.Context) (models.Blog, error) {
	title := ctx.QueryParam("keyword")
	blog, err := s.Database.GetBlogByTitle(title)
	if err != nil {
		return models.Blog{}, err
	}
	return blog, nil
}
