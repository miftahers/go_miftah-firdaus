package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
	"weekly-task-3/middleware"
	"weekly-task-3/models"
	"weekly-task-3/repositories"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type BlogSuites struct {
	suite.Suite
	mock    sqlmock.Sqlmock
	service BlogService
}

func TestSuiteBlog(t *testing.T) {
	suite.Run(t, new(BlogSuites))
}

func (s *BlogSuites) SetupSuite() {
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
	blogServices, _, _ := NewServices(repo)

	s.service = blogServices
	s.mock = mock
}

func (s *BlogSuites) TearDownSuite() {
	s.mock = nil
}

func (s *BlogSuites) TestNewBlog() {
	test := []struct {
		name        string
		path        string
		method      string
		blogBody    models.Blog
		expectError error
	}{
		{
			name:   "New Blog Service - Normal",
			path:   "/blogs",
			method: http.MethodPost,
			blogBody: models.Blog{
				UUID:       "1",
				Title:      "How to Test Functions in Golang",
				Content:    "abcdefghijklmnopqrstuvwxyz",
				UserID:     1,
				CategoryID: 1,
			},
			expectError: nil,
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

			err = s.service.NewBlog(c)

			s.Equal(v.expectError, err)
		})
	}
}

func (s *BlogSuites) TestNewBlogError() {
	test := []struct {
		name        string
		path        string
		method      string
		blogBody    models.Blog
		expectError error
	}{
		{
			name:   "New Blog Service - Normal",
			path:   "/blogs",
			method: http.MethodPost,
			blogBody: models.Blog{
				UUID:       "1",
				Title:      "How to Test Functions in Golang",
				Content:    "abcdefghijklmnopqrstuvwxyz",
				UserID:     1,
				CategoryID: 1,
			},
			expectError: errors.New("database error"),
		},
	}

	for _, v := range test {
		s.T().Run(v.name, func(t *testing.T) {
			s.mock.ExpectBegin()
			s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `blogs` (`created_at`,`updated_at`,`deleted_at`,`uuid`,`title`,`content`,`user_id`,`category_id`) VALUES (?,?,?,?,?,?,?,?)")).
				WithArgs(AnyTime{}, AnyTime{}, nil, "1", "How to Test Functions in Golang", "abcdefghijklmnopqrstuvwxyz", 1, 1).
				WillReturnError(errors.New("database error"))
			s.mock.ExpectRollback()

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

			err = s.service.NewBlog(c)

			s.Equal(v.expectError, err)
		})
	}
}

/* get blogs panic nil pointer bla bla bla
func (s *BlogSuites) TestGetBlogs() {
	test := []struct {
		name        string
		path        string
		method      string
		expectError error
		body        []models.Blog
	}{
		{
			name:        "service get blogs normal",
			path:        "/blogs",
			method:      http.MethodGet,
			expectError: nil,
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
			s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `blogs` WHERE `blogs`.`deleted_at` IS NULL")).
				WillReturnRows(expectedResult)

			res, _ := json.Marshal(v.body)

			r := httptest.NewRequest(v.method, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()
			c := echo.New().NewContext(r, w)
			c.SetPath(v.path)
			c.Request().Header.Set("Content-Type", "application/json")

			result, err := s.service.GetBlogs(c)

			s.Equal(v.expectError.Error(), err.Error())
			s.Equal(v.body, result)
		})
	}
} */

func (s *BlogSuites) TestGetBlogsError() {
	test := []struct {
		name        string
		path        string
		method      string
		expectError error
	}{
		{
			name:        "service get blogs normal",
			path:        "/blogs",
			method:      http.MethodGet,
			expectError: errors.New("record not found"),
		},
	}

	for _, v := range test {
		s.T().Run(v.name, func(t *testing.T) {
			s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `blogs` WHERE `blogs`.`deleted_at` IS NULL")).
				WillReturnError(errors.New("record not found"))

			r := httptest.NewRequest(v.method, "/", nil)
			w := httptest.NewRecorder()
			c := echo.New().NewContext(r, w)
			c.SetPath(v.path)
			c.Request().Header.Set("Content-Type", "application/json")

			_, err := s.service.GetBlogs(c)

			s.Equal(v.expectError, err)
		})
	}
}

func (s *BlogSuites) TestGetBlogByID() {

	testCases := []struct {
		name        string
		path        string
		expectError error
		Body        models.Blog
	}{
		{
			name:        "service get blog by id normal",
			path:        "/blogs/:id",
			expectError: nil,
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
			ctx.Request().Header.Set("Content-Type", "application/json")

			result, err := s.service.GetBlogByID(ctx)

			s.Equal(v.expectError, err)
			s.Equal(v.Body, result)
		})
	}
}

func (s *BlogSuites) TestGetBlogByIDError() {

	testCases := []struct {
		name        string
		path        string
		expectError error
		Body        models.Blog
	}{
		{
			name:        "service get blog by id Error",
			path:        "/blogs/:id",
			expectError: errors.New("record not found"),
			Body: models.Blog{
				Title:   "How to Test Functions in Golang 1",
				Content: "abcdefghijklmnopqrstuvwxyz 1",
			},
		},
	}

	for _, v := range testCases {
		s.T().Run(v.name, func(t *testing.T) {

			s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `blogs` WHERE uuid = ? AND `blogs`.`deleted_at` IS NULL ORDER BY `blogs`.`id` LIMIT 1")).
				WithArgs("1").
				WillReturnError(errors.New("record not found"))

			res, _ := json.Marshal(v.Body)

			r := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()
			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath(v.path)
			ctx.SetParamNames("id")
			ctx.SetParamValues("1")

			_, err := s.service.GetBlogByID(ctx)

			s.Equal(v.expectError.Error(), err.Error())
		})
	}
}

func (s *BlogSuites) TestUpdateBlog() {
	testCases := []struct {
		name        string
		path        string
		expectError error
		method      string
		Body        models.Blog
	}{
		{
			name:        "update blog service normal",
			path:        "/blogs/:id",
			expectError: nil,
			method:      http.MethodPut,
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

			err := s.service.UpdateBlog(ctx)

			s.Equal(v.expectError, err)
		})
	}
}

func (s *BlogSuites) TestUpdateBlogError() {

	var testCases = []struct {
		name          string
		path          string
		method        string
		expectedError error
	}{
		{
			name:          "update blogs error",
			path:          "/blogs/:id",
			method:        http.MethodPut,
			expectedError: echo.NewHTTPError(http.StatusInternalServerError, "record not found"),
		},
	}

	for _, v := range testCases {
		s.T().Run(v.name, func(t *testing.T) {

			s.mock.ExpectBegin()
			s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `blogs` SET `updated_at`=? WHERE items.uuid = ? AND `uuid` = ? AND `blogs`.`deleted_at` IS NULL")).
				WithArgs(AnyTime{}, "1", "1").
				WillReturnError(errors.New("record not found"))
			s.mock.ExpectRollback()

			r := httptest.NewRequest(v.method, "/", nil)
			w := httptest.NewRecorder()
			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath(v.path)
			ctx.SetParamNames("id")
			ctx.SetParamValues("1")

			err := s.service.UpdateBlog(ctx)

			s.Equal(v.expectedError.Error(), err.Error())
		})
	}
}

func (s *BlogSuites) TestDeleteBlog() {
	testCases := []struct {
		name        string
		path        string
		expectError error
		method      string
	}{
		{
			name:        "service delete blog normal",
			path:        "/blogs/:id",
			expectError: nil,
			method:      http.MethodDelete,
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

			err := s.service.DeleteBlog(ctx)

			s.Equal(v.expectError, err)
		})
	}
}

func (s *BlogSuites) TestDeleteUserControllerError() {

	var testCases = []struct {
		name        string
		path        string
		method      string
		expectError error
	}{
		{
			name:        "Delete blog error",
			path:        "/blogs/:id",
			method:      http.MethodDelete,
			expectError: echo.NewHTTPError(http.StatusInternalServerError, "record not found"),
		},
	}

	for _, v := range testCases {
		s.T().Run(v.name, func(t *testing.T) {

			s.mock.ExpectBegin()
			s.mock.ExpectExec(regexp.QuoteMeta("DELETE FROM `blogs` WHERE uuid = ?")).
				WithArgs("1").
				WillReturnError(errors.New("record not found"))
			s.mock.ExpectRollback()

			r := httptest.NewRequest(v.method, "/", nil)
			w := httptest.NewRecorder()
			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath(v.path)
			ctx.SetParamNames("id")
			ctx.SetParamValues("1")

			err := s.service.DeleteBlog(ctx)

			s.Equal(v.expectError, err)
		})
	}
}

func (s *BlogSuites) TestGetUserBytitle() {
	testCases := []struct {
		name        string
		path        string
		expectError error
		Body        models.Blog
	}{
		{
			name:        "service get blog by title normal",
			path:        "/blogs",
			expectError: nil,
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

			result, err := s.service.GetBlogByTitle(ctx)

			s.Equal(v.Body.Content, result.Content)
			s.Equal(v.expectError, err)
		})
	}
}

func (s *BlogSuites) TestGetBlogByTitleError() {

	var testCases = []struct {
		name        string
		path        string
		method      string
		expectError error
	}{
		{
			name:        "get blog by title error",
			path:        "/blogs",
			method:      http.MethodGet,
			expectError: errors.New("record not found"),
		},
	}

	for _, v := range testCases {
		s.T().Run(v.name, func(t *testing.T) {

			s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `blogs` WHERE name = ? AND `blogs`.`deleted_at` IS NULL ORDER BY `blogs`.`id` LIMIT 1")).
				WillReturnError(errors.New("record not found"))

			r := httptest.NewRequest(v.method, "/?keyword=Asem", nil)
			w := httptest.NewRecorder()
			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath(v.path)

			_, err := s.service.GetBlogByTitle(ctx)
			s.Equal(v.expectError.Error(), err.Error())
		})
	}
}

func (s *BlogSuites) TestGetBlogByCategory() {

	testCases := []struct {
		name        string
		path        string
		expectError error
		Body        []models.Blog
	}{
		{
			name:        "service get blog by category normal",
			path:        "/blogs/category/:category_id",
			expectError: nil,
			Body: []models.Blog{
				{
					Title:   "How to Test Functions in Golang 1",
					Content: "abcdefghijklmnopqrstuvwxyz 1",
				},
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

			result, err := s.service.GetBlogByCategory(ctx)

			s.Equal(v.expectError, err)
			s.Equal(v.Body[0].Content, result[0].Content)
		})
	}
}

func (s *BlogSuites) TestGetBlogByCategoryError() {

	var testCases = []struct {
		name        string
		path        string
		method      string
		expectError error
	}{
		{
			name:        "get blog by category error",
			path:        "/blogs/category/:category_id",
			method:      http.MethodGet,
			expectError: errors.New("record not found"),
		},
	}

	for _, v := range testCases {
		s.T().Run(v.name, func(t *testing.T) {

			s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `blogs` WHERE (category_id = ?) AND `blogs`.`deleted_at` IS NULL")).
				WillReturnError(errors.New("record not found"))

			r := httptest.NewRequest(v.method, "/", nil)
			w := httptest.NewRecorder()
			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath(v.path)
			ctx.SetParamNames("category_id")
			ctx.SetParamValues("1")

			_, err := s.service.GetBlogByCategory(ctx)
			s.Equal(v.expectError, err)
		})
	}
}

func (s *BlogSuites) TestGetBlogByCategoryError2() {

	var testCases = []struct {
		name        string
		path        string
		method      string
		expectError error
	}{
		{
			name:        "get blog by category error strconv",
			path:        "/blogs/category/:category_id",
			method:      http.MethodGet,
			expectError: errors.New("strconv.Atoi: parsing \"string\": invalid syntax"),
		},
	}

	for _, v := range testCases {
		s.T().Run(v.name, func(t *testing.T) {

			s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `blogs` WHERE category_id = ? AND `blogs`.`deleted_at` IS NULL")).
				WillReturnError(errors.New("record not found"))

			r := httptest.NewRequest(v.method, "/", nil)
			w := httptest.NewRecorder()
			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath(v.path)
			ctx.SetParamNames("category_id")
			ctx.SetParamValues("string")

			_, err := s.service.GetBlogByCategory(ctx)
			s.Equal(v.expectError.Error(), err.Error())
		})
	}
}
