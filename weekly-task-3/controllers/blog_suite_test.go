package controller

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
	"time"
	"weekly-task-3/middleware"
	"weekly-task-3/models"
	"weekly-task-3/repositories"
	"weekly-task-3/services"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type BlogSuites struct {
	suite.Suite
	mock    sqlmock.Sqlmock
	handler BlogHandler
}

type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func TestSuiteBlog(t *testing.T) {
	suite.Run(t, new(BlogSuites))
}

func (s *BlogSuites) SetupSuite() {
	// Create mock db
	db, mock, err := sqlmock.New()
	s.NoError(err)

	var GormDB *gorm.DB
	GormDB, err = gorm.Open(mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
		Conn:                      db,
	}), &gorm.Config{})
	s.NoError(err)

	repo := repositories.NewGorm(GormDB)
	blogServices, _, _ := services.NewServices(repo)
	handler := BlogHandler{
		BlogService: blogServices,
	}

	s.handler = handler
	s.mock = mock
}

func (s *BlogSuites) TearDownSuite() {
	s.mock = nil

}

func (s *BlogSuites) TestNewBlog() {
	test := []struct {
		name          string
		path          string
		method        string
		expectCode    int
		blogBody      models.Blog
		expectMessage string
	}{
		{
			name:       "New Blog - Normal",
			path:       "/blogs",
			method:     http.MethodPost,
			expectCode: http.StatusCreated,
			blogBody: models.Blog{
				UUID:       "1",
				Title:      "How to Test Functions in Golang",
				Content:    "abcdefghijklmnopqrstuvwxyz",
				UserID:     1,
				CategoryID: 1,
			},
			expectMessage: "created",
		},
	}

	for _, v := range test {
		s.T().Run(v.name, func(t *testing.T) {
			s.mock.ExpectBegin()
			s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `blogs` (`created_at`,`updated_at`,`deleted_at`,`uuid`,`title`,`content`,`user_id`,`category_id`) VALUES (?,?,?,?,?,?,?,?)")).
				WithArgs(AnyTime{}, AnyTime{}, nil, "1", "How to Test Functions in Golang", "abcdefghijklmnopqrstuvwxyz", 1, 1).
				WillReturnResult(sqlmock.NewResult(1, 1))
			s.mock.ExpectCommit()

			res, _ := json.Marshal(v.blogBody)
			r := httptest.NewRequest(v.method, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()
			c := echo.New().NewContext(r, w)
			c.SetPath(v.path)
			token, err := middleware.GetToken(1, "")
			if err != nil {
				s.Error(err)
			}
			value := fmt.Sprintf("Bearer %s", token)
			c.Request().Header.Add("Authorization", value)
			c.Request().Header.Set("Content-Type", "application/json")

			if s.NoError(s.handler.NewBlog(c)) {
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

func (s *BlogSuites) TestNewBlogError() {
	test := []struct {
		name        string
		path        string
		method      string
		expectError string
		blogBody    models.Blog
	}{
		{
			name:        "New Blog - Error",
			path:        "/blogs",
			method:      http.MethodPost,
			expectError: echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error").Error(),
			blogBody: models.Blog{
				Title:      "How to Test Functions in Golang",
				Content:    "abcdefghijklmnopqrstuvwxyz",
				UserID:     1,
				CategoryID: 1,
			},
		},
	}

	for _, v := range test {
		s.T().Run(v.name, func(t *testing.T) {
			s.mock.ExpectBegin()
			s.mock.ExpectExec(regexp.QuoteMeta("")).
				WithArgs().
				WillReturnError(errors.New("Internal Server Error"))
			s.mock.ExpectRollback()

			res, _ := json.Marshal(v.blogBody)
			r := httptest.NewRequest(v.method, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()
			c := echo.New().NewContext(r, w)
			c.SetPath(v.path)
			token, err := middleware.GetToken(0, "")
			if err != nil {
				s.Error(err)
			}
			value := fmt.Sprintf("Bearer %s", token)
			c.Request().Header.Add("Authorization", value)
			c.Request().Header.Set("Content-Type", "application/json")

			err = s.handler.NewBlog(c)

			s.Equal(v.expectError, err.Error())
		})
	}
}

func (s *BlogSuites) TestGetBlogs() {
	test := []struct {
		name       string
		path       string
		method     string
		expectCode int
		sizeData   int
		message    string
		body       []models.Blog
	}{
		{
			name:       "get blogs normal",
			path:       "/blogs",
			method:     http.MethodGet,
			expectCode: http.StatusOK,
			sizeData:   2,
			message:    "success",
			body: []models.Blog{
				{
					Title:   "How to Test Functions in Golang 1",
					Content: "abcdefghijklmnopqrstuvwxyz 1",
				},
				{
					Title:   "How to Test Functions in Golang 2",
					Content: "abcdefghijklmnopqrstuvwxyz 2",
				},
			},
		},
	}

	for _, v := range test {
		s.T().Run(v.name, func(t *testing.T) {
			expectedResult := sqlmock.NewRows([]string{"title", "content"}).
				AddRow("How to Test Functions in Golang 1", "abcdefghijklmnopqrstuvwxyz 1").
				AddRow("How to Test Functions in Golang 2", "abcdefghijklmnopqrstuvwxyz 2")
			s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `blogs`")).
				WillReturnRows(expectedResult)

			res, _ := json.Marshal(v.body)

			r := httptest.NewRequest(v.method, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()
			c := echo.New().NewContext(r, w)
			c.SetPath(v.path)
			c.Request().Header.Set("Content-Type", "application/json")

			if s.NoError(s.handler.GetBlogs(c)) {
				body := w.Body.Bytes()

				type response struct {
					Message string        `json:"message"`
					Data    []models.Blog `json:"data"`
				}
				var resp response
				err := json.Unmarshal(body, &resp)
				if err != nil {
					s.Error(err, "error unmarshalling")
				}

				s.Equal(v.expectCode, w.Result().StatusCode)
				s.Equal(v.sizeData, len(resp.Data))
				s.Equal(v.message, resp.Message)
				s.Equal(v.body, resp.Data)
			}
		})
	}
}

func (s *BlogSuites) TestGetBlogsError() {

	var testCases = []struct {
		name           string
		path           string
		method         string
		expectedResult string
	}{
		{
			name:           "get blogs error",
			path:           "/blogs",
			method:         http.MethodGet,
			expectedResult: echo.NewHTTPError(http.StatusInternalServerError, "Record Not found").Error(),
		},
	}

	for _, v := range testCases {
		s.T().Run(v.name, func(t *testing.T) {

			s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `blogs` WHERE `blogs`.`deleted_at` IS NULL")).
				WillReturnError(errors.New("Record Not found"))

			r := httptest.NewRequest(v.method, "/", nil)
			w := httptest.NewRecorder()
			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath(v.path)

			err := s.handler.GetBlogs(ctx)
			s.Equal(v.expectedResult, err.Error())
		})
	}
}

func (s *BlogSuites) TestGetBlogByID() {

	testCases := []struct {
		name       string
		path       string
		expectCode int
		message    string
		Body       models.Blog
	}{
		{
			name:       "get blog by id normal",
			path:       "/blogs/:id",
			expectCode: http.StatusOK,
			message:    "success",
			Body: models.Blog{
				Title:   "How to Test Functions in Golang 1",
				Content: "abcdefghijklmnopqrstuvwxyz 1",
			},
		},
	}

	for _, v := range testCases {
		s.T().Run(v.name, func(t *testing.T) {

			expectedResult := sqlmock.NewRows([]string{"title", "content"}).
				AddRow("How to Test Functions in Golang 1", "abcdefghijklmnopqrstuvwxyz 1")
			s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `blogs` WHERE uuid = ? AND `blogs`.`deleted_at` IS NULL ORDER BY `blogs`.`id` LIMIT 1")).
				WithArgs("1").
				WillReturnRows(expectedResult)

			res, _ := json.Marshal(v.Body)

			r := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()
			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath(v.path)
			ctx.SetParamNames("id")
			ctx.SetParamValues("1")

			if s.NoError(s.handler.GetBlogById(ctx)) {
				body := w.Body.Bytes()

				type response struct {
					Message string      `json:"message"`
					Data    models.Blog `json:"data"`
				}

				var resp response
				if err := json.Unmarshal(body, &resp); err != nil {
					s.Error(err, "error unmarshalling")
				}

				s.Equal(v.expectCode, w.Result().StatusCode)
				s.Equal(v.message, resp.Message)
				s.Equal(v.Body, resp.Data)
			}
		})
	}
}

func (s *BlogSuites) TestGetBlogByIDError() {

	var testCases = []struct {
		name           string
		path           string
		method         string
		expectedResult string
	}{
		{
			name:           "get blog by id error",
			path:           "/blogs/:id",
			method:         http.MethodGet,
			expectedResult: echo.NewHTTPError(http.StatusInternalServerError, "Record Not found").Error(),
		},
	}

	for _, v := range testCases {
		s.T().Run(v.name, func(t *testing.T) {

			s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `blogs` WHERE uuid = ? AND `blogs`.`deleted_at` IS NULL ORDER BY `blogs`.`id` LIMIT 1")).
				WillReturnError(errors.New("Record Not found"))

			r := httptest.NewRequest(v.method, "/", nil)
			w := httptest.NewRecorder()
			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath(v.path)
			ctx.SetParamNames("id")
			ctx.SetParamValues("1")

			err := s.handler.GetBlogById(ctx).Error()
			s.Equal(v.expectedResult, err)
		})
	}
}

func (s *BlogSuites) TestUpdateBlog() {
	testCases := []struct {
		name       string
		path       string
		expectCode int
		method     string
		message    string
		Body       models.Blog
	}{
		{
			name:       "update blog normal",
			path:       "/blogs/:id",
			expectCode: http.StatusOK,
			method:     http.MethodPut,
			message:    "updated",
			Body: models.Blog{
				Title:   "this is title",
				Content: "this is the content of <title>",
			},
		},
	}

	for _, v := range testCases {
		s.T().Run(v.name, func(t *testing.T) {
			s.mock.ExpectBegin()
			s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `blogs` SET `updated_at`=?,`title`=?,`content`=? WHERE items.uuid = ? AND `uuid` = ? AND `blogs`.`deleted_at` IS NULL")).
				WithArgs(AnyTime{}, "this is title", "this is the content of <title>", "1", "1").
				WillReturnResult(sqlmock.NewResult(1, 1))
			s.mock.ExpectCommit()

			res, _ := json.Marshal(v.Body)
			r := httptest.NewRequest(v.method, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()
			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath(v.path)
			ctx.SetParamNames("id")
			ctx.SetParamValues("1")
			ctx.Request().Header.Set("Content-Type", "application/json")

			if s.NoError(s.handler.UpdateBlog(ctx)) {
				body := w.Body.Bytes()

				type Response struct {
					Message string `json:"message"`
				}
				var resp Response
				if err := json.Unmarshal(body, &resp); err != nil {
					s.Error(err, "error unmarshalling")
				}

				s.Equal(v.expectCode, w.Result().StatusCode)
				s.Equal(v.message, resp.Message)
			}
		})
	}
}

func (s *BlogSuites) TestUpdateBlogError() {

	var testCases = []struct {
		name                  string
		path                  string
		method                string
		expectedResult        string
		expectedStrconvResult string
	}{
		{
			name:           "update blogs error",
			path:           "/blogs/:id",
			method:         http.MethodPut,
			expectedResult: echo.NewHTTPError(http.StatusInternalServerError, "code=500, message=Record Not found").Error(),
		},
	}

	for _, v := range testCases {
		s.T().Run(v.name, func(t *testing.T) {

			s.mock.ExpectBegin()
			s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `blogs` SET `updated_at`=? WHERE items.uuid = ? AND `uuid` = ? AND `blogs`.`deleted_at` IS NULL")).
				WithArgs(AnyTime{}, "1", "1").
				WillReturnError(errors.New("Record Not found"))
			s.mock.ExpectRollback()

			r := httptest.NewRequest(v.method, "/", nil)
			w := httptest.NewRecorder()
			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath(v.path)
			ctx.SetParamNames("id")
			ctx.SetParamValues("1")

			err := s.handler.UpdateBlog(ctx).Error()
			s.Equal(v.expectedResult, err)
		})
	}
}

func (s *BlogSuites) TestDeleteBlog() {
	testCases := []struct {
		name       string
		path       string
		expectCode int
		method     string
		message    string
	}{
		{
			name:       "delete blog normal",
			path:       "/blogs/:id",
			expectCode: http.StatusOK,
			method:     http.MethodDelete,
			message:    "deleted",
		},
	}

	for _, v := range testCases {
		s.T().Run(v.name, func(t *testing.T) {
			if v.method == http.MethodDelete {
				s.mock.ExpectBegin()
				s.mock.ExpectExec(regexp.QuoteMeta("DELETE FROM `blogs` WHERE uuid = ?")).
					WithArgs("1").
					WillReturnResult(sqlmock.NewResult(1, 1))
				s.mock.ExpectCommit()
			}

			r := httptest.NewRequest(v.method, "/", nil)
			w := httptest.NewRecorder()
			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath(v.path)
			ctx.SetParamNames("id")
			ctx.SetParamValues("1")

			if s.NoError(s.handler.DeleteBlog(ctx)) {
				body := w.Body.Bytes()

				type Response struct {
					Message string `json:"message"`
				}
				var resp Response
				if err := json.Unmarshal(body, &resp); err != nil {
					s.Error(err, "error unmarshalling")
				}

				s.Equal(v.expectCode, w.Result().StatusCode)
				s.Equal(v.message, resp.Message)
			}
		})
	}
}

func (s *BlogSuites) TestDeleteBlogError() {

	var testCases = []struct {
		name           string
		path           string
		method         string
		expectedResult string
	}{
		{
			name:           "Delete blog error",
			path:           "/blogs/:id",
			method:         http.MethodDelete,
			expectedResult: echo.NewHTTPError(http.StatusInternalServerError, "code=500, message=Record Not found").Error(),
		},
	}

	for _, v := range testCases {
		s.T().Run(v.name, func(t *testing.T) {

			s.mock.ExpectBegin()
			s.mock.ExpectExec(regexp.QuoteMeta("DELETE FROM `blogs` WHERE uuid = ?")).
				WithArgs("1").
				WillReturnError(errors.New("Record Not found"))
			s.mock.ExpectRollback()

			r := httptest.NewRequest(v.method, "/", nil)
			w := httptest.NewRecorder()
			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath(v.path)
			ctx.SetParamNames("id")
			ctx.SetParamValues("1")

			err := s.handler.DeleteBlog(ctx).Error()
			s.Equal(v.expectedResult, err)
		})
	}
}

func (s *BlogSuites) TestGetBlogBytitle() {
	testCases := []struct {
		name       string
		path       string
		expectCode int
		message    string
		Body       models.Blog
	}{
		{
			name:       "get blog by title normal",
			path:       "/blogs",
			expectCode: http.StatusOK,
			message:    "success",
			Body: models.Blog{
				Title:   "Apple",
				Content: "abcdefghijklmnopqrstuvwxyz 1",
			},
		},
	}

	for _, v := range testCases {
		s.T().Run(v.name, func(t *testing.T) {

			expectedResult := sqlmock.NewRows([]string{"title", "content"}).
				AddRow("Apple", "abcdefghijklmnopqrstuvwxyz 1")
			s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `blogs` WHERE name = ? AND `blogs`.`deleted_at` IS NULL ORDER BY `blogs`.`id` LIMIT 1")).
				WithArgs("Apple").
				WillReturnRows(expectedResult)

			res, _ := json.Marshal(v.Body)

			r := httptest.NewRequest(http.MethodGet, "/?keyword=Apple", bytes.NewBuffer(res))
			w := httptest.NewRecorder()
			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath(v.path)
			ctx.Request().Header.Set("Content-Type", "application/json")

			if s.NoError(s.handler.GetBlogs(ctx)) {
				body := w.Body.Bytes()

				type response struct {
					Message string      `json:"message"`
					Data    models.Blog `json:"data"`
				}

				var resp response
				if err := json.Unmarshal(body, &resp); err != nil {
					s.Error(err, "error unmarshalling")
				}

				s.Equal(v.expectCode, w.Result().StatusCode)
				s.Equal(v.message, resp.Message)
				s.Equal(v.Body, resp.Data)
			}
		})
	}
}

func (s *BlogSuites) TestGetBlogByTitleError() {

	var testCases = []struct {
		name           string
		path           string
		method         string
		expectedResult string
	}{
		{
			name:           "get blog by title error",
			path:           "/blogs",
			method:         http.MethodGet,
			expectedResult: echo.NewHTTPError(http.StatusInternalServerError, "Record Not found").Error(),
		},
	}

	for _, v := range testCases {
		s.T().Run(v.name, func(t *testing.T) {

			s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `blogs` WHERE name = ? AND `blogs`.`deleted_at` IS NULL ORDER BY `blogs`.`id` LIMIT 1")).
				WillReturnError(errors.New("Record Not found"))

			r := httptest.NewRequest(v.method, "/?keyword=Asem", nil)
			w := httptest.NewRecorder()
			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath(v.path)

			err := s.handler.GetBlogs(ctx).Error()
			s.Equal(v.expectedResult, err)
		})
	}
}

func (s *BlogSuites) TestGetBlogByCategory() {

	testCases := []struct {
		name       string
		path       string
		expectCode int
		message    string
		Body       []models.Blog
	}{
		{
			name:       "get blog by category normal",
			path:       "/blogs/category/:category_id",
			expectCode: http.StatusOK,
			message:    "success",
			Body: []models.Blog{
				{Title: "How to Test Functions in Golang 1",
					Content: "abcdefghijklmnopqrstuvwxyz 1"},
			},
		},
	}

	for _, v := range testCases {
		s.T().Run(v.name, func(t *testing.T) {

			expectedResult := sqlmock.NewRows([]string{"title", "content"}).
				AddRow("How to Test Functions in Golang 1", "abcdefghijklmnopqrstuvwxyz 1")
			s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `blogs` WHERE (category_id = ?) AND `blogs`.`deleted_at` IS NULL")).
				WithArgs(1).
				WillReturnRows(expectedResult)

			res, _ := json.Marshal(v.Body)

			r := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()
			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath(v.path)
			ctx.SetParamNames("category_id")
			ctx.SetParamValues("1")
			ctx.Request().Header.Set("Content-Type", "application/json")

			if s.NoError(s.handler.GetBlogByCategory(ctx)) {
				body := w.Body.Bytes()

				type response struct {
					Message string        `json:"message"`
					Data    []models.Blog `json:"data"`
				}

				var resp response
				if err := json.Unmarshal(body, &resp); err != nil {
					s.Error(err, "error unmarshalling")
				}

				s.Equal(v.expectCode, w.Result().StatusCode)
				s.Equal(v.message, resp.Message)
				s.Equal(v.Body, resp.Data)
			}
		})
	}
}

func (s *BlogSuites) TestGetBlogByCategoryError() {

	var testCases = []struct {
		name           string
		path           string
		method         string
		expectedResult string
	}{
		{
			name:           "get blog by category error",
			path:           "/blogs/category/:category_id",
			method:         http.MethodGet,
			expectedResult: echo.NewHTTPError(http.StatusInternalServerError, "Record Not found").Error(),
		},
	}

	for _, v := range testCases {
		s.T().Run(v.name, func(t *testing.T) {

			s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `blogs` WHERE (category_id = ?) AND `blogs`.`deleted_at` IS NULL")).
				WillReturnError(errors.New("Record Not found"))

			r := httptest.NewRequest(v.method, "/", nil)
			w := httptest.NewRecorder()
			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath(v.path)
			ctx.SetParamNames("category_id")
			ctx.SetParamValues("1")

			err := s.handler.GetBlogByCategory(ctx).Error()
			s.Equal(v.expectedResult, err)
		})
	}
}
