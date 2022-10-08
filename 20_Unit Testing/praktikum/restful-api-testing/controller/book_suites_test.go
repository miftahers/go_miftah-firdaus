package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"regexp"
	"restful-api-testing/config"
	"restful-api-testing/model"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type suiteBooks struct {
	suite.Suite
	mock sqlmock.Sqlmock
}

func TestSuiteBooks(t *testing.T) {
	suite.Run(t, new(suiteBooks))
}

func (s *suiteBooks) SetupSuite() {
	// Create mock db
	db, mock, err := sqlmock.New()
	s.NoError(err)

	var gormDB *gorm.DB
	gormDB, err = gorm.Open(mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
		Conn:                      db,
	}), &gorm.Config{})
	s.NoError(err)

	config.DB = gormDB
	s.mock = mock
	// End of create mock db
}

func (s *suiteBooks) TestGetBooksController() {

	var testCases = []struct {
		name       string
		path       string
		expectCode int
		sizeData   int
		message    string
		Body       []model.Book
	}{
		{
			name:       "get Books normal",
			path:       "/books",
			expectCode: http.StatusOK,
			sizeData:   2,
			message:    "success",
			Body: []model.Book{
				{
					Title:        "book1",
					Category:     "Category1",
					Release_Year: "2022",
					Writter:      "Writter1",
				},
				{
					Title:        "book2",
					Category:     "Category2",
					Release_Year: "2022",
					Writter:      "Writter2",
				},
			},
		},
	}

	for _, v := range testCases {
		s.T().Run(v.name, func(t *testing.T) {

			if v.expectCode == http.StatusOK {
				expectedResult := sqlmock.NewRows([]string{"title", "category", "release_year", "writter"}).
					AddRow("book1", "Category1", "2022", "Writter1").
					AddRow("book2", "Category2", "2022", "Writter2")
				s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `books` WHERE `books`.`deleted_at` IS NULL")).
					WillReturnRows(expectedResult)
			}

			res, _ := json.Marshal(v.Body)

			r := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()
			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath(v.path)

			if s.NoError(GetBooksController(ctx)) {
				body := w.Body.Bytes()

				type Response struct {
					Message string       `json:"message"`
					Data    []model.User `json:"data"`
				}
				var user Response
				if err := json.Unmarshal(body, &user); err != nil {
					s.Error(err, "error unmarshalling")
				}

				s.Equal(v.expectCode, w.Result().StatusCode)
				s.Equal(v.message, user.Message)
				s.Equal(v.sizeData, len(user.Data))
				// s.Equal(v.Body, user.Data)
			}
		})
	}
}
func (s *suiteBooks) TestGetBooksControllerInvalid() {

	var testCases = []struct {
		name           string
		path           string
		method         string
		expectedResult string
	}{
		{
			name:           "get books error",
			path:           "/books",
			method:         http.MethodGet,
			expectedResult: echo.NewHTTPError(http.StatusInternalServerError, "Record Not found").Error(),
		},
	}

	for _, v := range testCases {
		s.T().Run(v.name, func(t *testing.T) {

			s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `books` WHERE `books`.`deleted_at` IS NULL")).
				WillReturnError(errors.New("Record Not found"))

			r := httptest.NewRequest(v.method, "/", nil)
			w := httptest.NewRecorder()
			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath(v.path)

			err := GetBooksController(ctx).Error()
			s.Equal(v.expectedResult, err)
		})
	}
}

func (s *suiteBooks) TestGetBookController() {

	testCases := []struct {
		name       string
		path       string
		expectCode int
		message    string
		Body       model.Book
	}{
		{
			name:       "get user normal",
			path:       "/books/:id",
			expectCode: http.StatusOK,
			message:    "success",
			Body: model.Book{
				Title:        "book1",
				Category:     "Category1",
				Release_Year: "2022",
				Writter:      "Writter1",
			},
		},
	}

	for _, v := range testCases {
		s.T().Run(v.name, func(t *testing.T) {
			if v.expectCode == http.StatusOK {
				expectedResult := sqlmock.NewRows([]string{"title", "category", "release_year", "writter"}).
					AddRow("book1", "Category1", "2022", "Writter1")
				s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `books` WHERE id = ? AND `books`.`deleted_at` IS NULL ORDER BY `books`.`id` LIMIT 1")).
					WithArgs(1).
					WillReturnRows(expectedResult)
			}

			res, _ := json.Marshal(v.Body)

			r := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()
			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath(v.path)
			ctx.SetParamNames("id")
			ctx.SetParamValues("1")

			if s.NoError(GetBookController(ctx)) {
				body := w.Body.Bytes()

				type Response struct {
					Message string     `json:"message"`
					Data    model.User `json:"data"`
				}

				var user Response
				if err := json.Unmarshal(body, &user); err != nil {
					s.Error(err, "error unmarshalling")
				}

				s.Equal(v.expectCode, w.Result().StatusCode)
				s.Equal(v.message, user.Message)
				// s.Equal(v.Body, user.Data)
			}
		})
	}
}
func (s *suiteBooks) TestGetBookControllerInvalid() {

	var testCases = []struct {
		name           string
		path           string
		method         string
		expectedResult string
	}{
		{
			name:           "get book error",
			path:           "/books/:id",
			method:         http.MethodGet,
			expectedResult: echo.NewHTTPError(http.StatusInternalServerError, "Record Not found").Error(),
		},
	}

	for _, v := range testCases {
		s.T().Run(v.name, func(t *testing.T) {

			s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `books` WHERE id = ? AND `books`.`deleted_at` IS NULL ORDER BY `books`.`id` LIMIT 1")).
				WillReturnError(errors.New("Record Not found"))

			r := httptest.NewRequest(v.method, "/", nil)
			w := httptest.NewRecorder()
			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath(v.path)
			ctx.SetParamNames("id")
			ctx.SetParamValues("1")

			err := GetBookController(ctx).Error()
			s.Equal(v.expectedResult, err)
		})
	}
}

func (s *suiteBooks) TestCreateBookController() {

	var testCases = []struct {
		name       string
		path       string
		expectCode int
		method     string
		message    string
		Body       model.Book
	}{
		{
			name:       "create book normal",
			path:       "/books",
			expectCode: http.StatusCreated,
			method:     http.MethodPost,
			message:    "book created!",
			Body:       model.Book{},
		},
	}

	for _, v := range testCases {
		s.T().Run(v.name, func(t *testing.T) {

			if v.expectCode == http.StatusCreated {
				s.mock.ExpectBegin()
				s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `books` (`created_at`,`updated_at`,`deleted_at`,`title`,`category`,`release_year`,`writter`) VALUES (?,?,?,?,?,?,?)")).
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), nil, "", "", "", "").
					WillReturnResult(sqlmock.NewResult(1, 1))
				s.mock.ExpectCommit()
			}

			res, _ := json.Marshal(v.Body)
			r := httptest.NewRequest(v.method, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()
			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath(v.path)

			if s.NoError(CreateBookController(ctx)) {
				body := w.Body.Bytes()

				type Response struct {
					Message string `json:"message"`
				}
				var user Response
				if err := json.Unmarshal(body, &user); err != nil {
					s.Error(err, "error unmarshalling")
				}

				s.Equal(v.expectCode, w.Result().StatusCode)
				s.Equal(v.message, user.Message)
			}

		})
	}
}
func (s *suiteBooks) TestCreateBookControllerInvalid() {

	var testCases = []struct {
		name           string
		path           string
		method         string
		expectedResult string
	}{
		{
			name:           "create book error",
			path:           "/books",
			method:         http.MethodPost,
			expectedResult: echo.NewHTTPError(http.StatusInternalServerError, "Record Not found").Error(),
		},
	}

	for _, v := range testCases {
		s.T().Run(v.name, func(t *testing.T) {

			s.mock.ExpectBegin()
			s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `books` (`created_at`,`updated_at`,`deleted_at`,`title`,`category`,`release_year`,`writter`) VALUES (?,?,?,?,?,?,?)")).
				WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), nil, "", "", "", "").
				WillReturnError(errors.New("Record Not found"))
			s.mock.ExpectRollback()

			r := httptest.NewRequest(v.method, "/", nil)
			w := httptest.NewRecorder()
			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath(v.path)

			err := CreateBookController(ctx).Error()
			s.Equal(v.expectedResult, err)
		})
	}
}

func (s *suiteBooks) TestUpdateBookController() {
	testCases := []struct {
		name       string
		path       string
		expectCode int
		method     string
		message    string
		Body       model.Book
	}{
		{
			name:       "update book normal",
			path:       "/book/:id",
			expectCode: http.StatusOK,
			method:     http.MethodPut,
			message:    "book updated!",
			Body:       model.Book{},
		},
	}

	for _, v := range testCases {
		s.T().Run(v.name, func(t *testing.T) {
			if v.expectCode == http.StatusOK {
				s.mock.ExpectBegin()
				s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `books` SET `updated_at`=? WHERE id = ? AND `books`.`deleted_at` IS NULL")).
					WithArgs(sqlmock.AnyArg(), 1).
					WillReturnResult(sqlmock.NewResult(1, 1))
				s.mock.ExpectCommit()
			}

			res, _ := json.Marshal(v.Body)
			r := httptest.NewRequest(v.method, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()
			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath(v.path)
			ctx.SetParamNames("id")
			ctx.SetParamValues("1")

			if s.NoError(UpdateBookController(ctx)) {
				body := w.Body.Bytes()

				type Response struct {
					Message string `json:"message"`
				}
				var user Response
				if err := json.Unmarshal(body, &user); err != nil {
					s.Error(err, "error unmarshalling")
				}

				s.Equal(v.expectCode, w.Result().StatusCode)
				s.Equal(v.message, user.Message)
			}
		})
	}
}
func (s *suiteUsers) TestUpdateBookControllerInvalid() {

	var testCases = []struct {
		name                  string
		path                  string
		method                string
		expectedResult        string
		expectedStrconvResult string
	}{
		{
			name:           "update book error",
			path:           "/book/:id",
			method:         http.MethodPut,
			expectedResult: echo.NewHTTPError(http.StatusInternalServerError, "Record Not found").Error(),
		},
	}

	for _, v := range testCases {
		s.T().Run(v.name, func(t *testing.T) {

			s.mock.ExpectBegin()
			s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `books` SET `updated_at`=? WHERE id = ? AND `books`.`deleted_at` IS NULL")).
				WithArgs(sqlmock.AnyArg(), 1).
				WillReturnError(errors.New("Record Not found"))
			s.mock.ExpectRollback()

			r := httptest.NewRequest(v.method, "/", nil)
			w := httptest.NewRecorder()
			e := echo.New()
			ctx1 := e.NewContext(r, w)
			ctx1.SetPath(v.path)
			ctx1.SetParamNames("id")
			ctx1.SetParamValues("1")

			err := UpdateBookController(ctx1).Error()
			s.Equal(v.expectedResult, err)
		})
	}
}

func (s *suiteBooks) TestDeleteBookController() {
	testCases := []struct {
		name       string
		path       string
		expectCode int
		method     string
		message    string
	}{
		{
			name:       "delete book normal",
			path:       "/books/:id",
			expectCode: http.StatusOK,
			method:     http.MethodDelete,
			message:    "book deleted",
		},
	}

	for _, v := range testCases {
		s.T().Run(v.name, func(t *testing.T) {
			if v.method == http.MethodDelete {
				s.mock.ExpectBegin()
				s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `books` SET `deleted_at`=? WHERE `books`.`id` = ? AND `books`.`deleted_at` IS NULL")).
					WithArgs(sqlmock.AnyArg(), 1).
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

			if s.NoError(DeleteBookController(ctx)) {
				body := w.Body.Bytes()

				type Response struct {
					Message string `json:"message"`
				}
				var user Response
				if err := json.Unmarshal(body, &user); err != nil {
					s.Error(err, "error unmarshalling")
				}

				s.Equal(v.expectCode, w.Result().StatusCode)
				s.Equal(v.message, user.Message)
			}
		})
	}
}
func (s *suiteBooks) TestDeleteBookControllerInvalid() {

	var testCases = []struct {
		name           string
		path           string
		method         string
		expectedResult string
	}{
		{
			name:           "Delete user error",
			path:           "/users/:id",
			method:         http.MethodDelete,
			expectedResult: echo.NewHTTPError(http.StatusInternalServerError, "Record Not found").Error(),
		},
	}

	for _, v := range testCases {
		s.T().Run(v.name, func(t *testing.T) {

			s.mock.ExpectBegin()
			s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `books` SET `deleted_at`=? WHERE `books`.`id` = ? AND `books`.`deleted_at` IS NULL")).
				WithArgs(sqlmock.AnyArg(), 1).
				WillReturnError(errors.New("Record Not found"))
			s.mock.ExpectRollback()

			r := httptest.NewRequest(v.method, "/", nil)
			w := httptest.NewRecorder()
			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath(v.path)
			ctx.SetParamNames("id")
			ctx.SetParamValues("1")

			err := DeleteBookController(ctx).Error()
			s.Equal(v.expectedResult, err)
		})
	}
}

func (s *suiteBooks) TearDownSuite() {
	config.DB = nil
	s.mock = nil
}
