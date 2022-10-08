package controller

import (
	"net/http"
	cfg "restful-api-testing/config"
	"restful-api-testing/model"
	"strconv"

	"github.com/labstack/echo"
)

func GetBooksController(ctx echo.Context) error {
	var books []model.Book
	if err := cfg.DB.Find(&books).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    books,
	})
}

func GetBookController(ctx echo.Context) error {
	var book []model.Book
	id, _ := strconv.Atoi(ctx.Param("id"))

	if err := cfg.DB.First(&book, "id = ?", id).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    book,
	})
}

func CreateBookController(ctx echo.Context) error {

	book := model.Book{}

	ctx.Bind(&book)

	if err := cfg.DB.Save(&book).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, map[string]interface{}{
		"message": "book created!",
	})
}

func UpdateBookController(ctx echo.Context) error {

	var book model.Book

	id, _ := strconv.Atoi(ctx.Param("id"))

	title := ctx.FormValue("title")
	category := ctx.FormValue("category")
	release_year := ctx.FormValue("release_year")
	writter := ctx.FormValue("writter")

	updatedBook := model.Book{
		Title:        title,
		Category:     category,
		Release_Year: release_year,
		Writter:      writter,
	}
	if err := cfg.DB.Model(&book).Where("id = ?", id).Updates(updatedBook).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "book updated!",
	})
}

func DeleteBookController(ctx echo.Context) error {

	var book model.Book

	id, _ := strconv.Atoi(ctx.Param("id"))

	if err := cfg.DB.Delete(&book, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "book deleted",
	})
}
