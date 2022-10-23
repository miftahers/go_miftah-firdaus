package controller

import (
	"net/http"
	"weekly-task-3/services"

	"github.com/labstack/echo"
)

type BlogHandler struct {
	services.BlogService
}

func (h *BlogHandler) GetBlogs(ctx echo.Context) error {
	if ctx.QueryParam("keyword") == "" {
		result, err := h.BlogService.GetBlogs(ctx)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"message": "success",
			"data":    result,
		})
	} else {
		result, err := h.BlogService.GetBlogByTitle(ctx)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"message": "success",
			"data":    result,
		})
	}
}

func (h *BlogHandler) GetBlogById(ctx echo.Context) error {
	blog, err := h.BlogService.GetBlogByID(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    blog,
	})
}

func (h *BlogHandler) NewBlog(ctx echo.Context) error {
	err := h.BlogService.NewBlog(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusCreated, map[string]interface{}{
		"message": "created",
	})
}

func (h *BlogHandler) UpdateBlog(ctx echo.Context) error {
	err := h.BlogService.UpdateBlog(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "updated",
	})
}

func (h *BlogHandler) DeleteBlog(ctx echo.Context) error {
	err := h.BlogService.DeleteBlog(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "deleted",
	})
}

func (h *BlogHandler) GetBlogByCategory(ctx echo.Context) error {
	blogs, err := h.BlogService.GetBlogByCategory(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    blogs,
	})
}
