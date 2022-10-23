package services

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
	"time"
	"weekly-task-3/models"
	"weekly-task-3/repositories"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}
func TestNewCategory(t *testing.T) {
	test := []struct {
		name         string
		expectResult error
		body         models.Category
		method       string
		path         string
	}{
		{
			name:         "new category service normal",
			expectResult: nil,
			body: models.Category{
				Name: "Makanan",
			},
			method: http.MethodPost,
			path:   "/blogs/category",
		},
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}

	var gormDB *gorm.DB
	gormDB, err = gorm.Open(mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
		Conn:                      db,
	}), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	repo := repositories.NewGorm(gormDB)
	_, _, service := NewServices(repo)
	for _, v := range test {
		t.Run(v.name, func(t *testing.T) {

			mock.ExpectBegin()
			mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `categories` (`created_at`,`updated_at`,`deleted_at`,`name`) VALUES (?,?,?,?)")).
				WithArgs(AnyTime{}, AnyTime{}, nil, "Makanan").
				WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectCommit()

			res, _ := json.Marshal(v.body)

			r := httptest.NewRequest(v.method, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()

			c := echo.New().NewContext(r, w)
			c.SetPath(v.path)
			c.Request().Header.Set("Content-Type", "application/json")

			err := service.NewCategory(c)
			if err != nil {
				t.Error(err)
			}
		})
	}
}

func TestNewCategoryError(t *testing.T) {
	test := []struct {
		name         string
		expectResult error
		body         models.Category
		method       string
		path         string
	}{
		{
			name:         "new category service error",
			expectResult: errors.New("Internal Server Error"),
			body: models.Category{
				Name: "Makanan",
			},
			method: http.MethodPost,
			path:   "/blogs/category",
		},
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}

	var gormDB *gorm.DB
	gormDB, err = gorm.Open(mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
		Conn:                      db,
	}), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	repo := repositories.NewGorm(gormDB)
	_, _, service := NewServices(repo)
	for _, v := range test {
		t.Run(v.name, func(t *testing.T) {

			mock.ExpectBegin()
			mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `categories` (`created_at`,`updated_at`,`deleted_at`,`name`) VALUES (?,?,?,?)")).
				WithArgs(AnyTime{}, AnyTime{}, nil, "Makanan").
				WillReturnError(errors.New("Internal Server Error"))
			mock.ExpectRollback()

			res, _ := json.Marshal(v.body)

			r := httptest.NewRequest(v.method, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()

			c := echo.New().NewContext(r, w)
			c.SetPath(v.path)
			c.Request().Header.Set("Content-Type", "application/json")

			err := service.NewCategory(c)

			if err.Error() != v.expectResult.Error() {
				t.Error(err)
			}
		})
	}
}
