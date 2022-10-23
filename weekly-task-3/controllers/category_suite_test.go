package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
	"weekly-task-3/models"
	"weekly-task-3/repositories"
	"weekly-task-3/services"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type CategorySuites struct {
	suite.Suite
	mock    sqlmock.Sqlmock
	handler CategoryHandler
}

func TestSuiteCategory(t *testing.T) {
	suite.Run(t, new(CategorySuites))
}

func (s *CategorySuites) SetupSuite() {
	// Create mock db
	db, mock, err := sqlmock.New()
	s.NoError(err)

	var gormDB *gorm.DB
	gormDB, err = gorm.Open(mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
		Conn:                      db,
	}), &gorm.Config{})
	s.NoError(err)

	repo := repositories.NewGorm(gormDB)
	_, _, categoryServices := services.NewServices(repo)
	handler := CategoryHandler{
		CategoryService: categoryServices,
	}

	s.handler = handler
	s.mock = mock
}

func (s *CategorySuites) TearDownSuite() {
	s.mock = nil
}

func (s *CategorySuites) TestNewCategory() {
	test := []struct {
		name          string
		path          string
		method        string
		expectCode    int
		body          models.Category
		expectMessage string
	}{
		{
			name:       "Create Category - Normal",
			path:       "/blogs/category",
			method:     http.MethodPost,
			expectCode: http.StatusCreated,
			body: models.Category{
				Name: "Politik",
			},
			expectMessage: "created",
		},
	}

	for _, v := range test {
		s.T().Run(v.name, func(t *testing.T) {
			s.mock.ExpectBegin()
			s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `categories` (`created_at`,`updated_at`,`deleted_at`,`name`) VALUES (?,?,?,?)")).
				WithArgs(AnyTime{}, AnyTime{}, nil, "Politik").
				WillReturnResult(sqlmock.NewResult(1, 1))
			s.mock.ExpectCommit()

			res, _ := json.Marshal(v.body)
			r := httptest.NewRequest(v.method, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()
			c := echo.New().NewContext(r, w)
			c.SetPath(v.path)
			c.Request().Header.Set("Content-Type", "application/json")

			if s.NoError(s.handler.NewCategory(c)) {
				body := w.Body.Bytes()

				type Response struct {
					Message string `json:"message"`
				}
				var resp Response
				err := json.Unmarshal(body, &resp)
				if err != nil {
					s.Error(err, "error unmarshalling")
				}

				s.Equal(v.expectCode, w.Result().StatusCode)
				s.Equal(v.expectMessage, resp.Message)
			}
		})
	}
}

func (s *CategorySuites) TestNewCategoryError() {
	test := []struct {
		name        string
		path        string
		method      string
		expectError string
		body        models.Category
	}{
		{
			name:        "New Category - Error",
			path:        "/blogs/category",
			method:      http.MethodPost,
			expectError: echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error").Error(),
			body: models.Category{
				Name: "Politik",
			},
		},
	}

	for _, v := range test {
		s.T().Run(v.name, func(t *testing.T) {
			s.mock.ExpectBegin()
			s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `categories` (`created_at`,`updated_at`,`deleted_at`,`name`) VALUES (?,?,?,?)")).
				WithArgs().
				WillReturnError(errors.New("Internal Server Error"))
			s.mock.ExpectRollback()

			res, _ := json.Marshal(v.body)
			r := httptest.NewRequest(v.method, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()
			c := echo.New().NewContext(r, w)
			c.SetPath(v.path)
			c.Request().Header.Set("Content-Type", "application/json")

			err := s.handler.NewCategory(c)

			s.Equal(v.expectError, err.Error())
		})
	}
}
