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

type suiteUsers struct {
	suite.Suite
	mock sqlmock.Sqlmock
}

func TestSuiteUsers(t *testing.T) {
	suite.Run(t, new(suiteUsers))
}
func (s *suiteUsers) SetupSuite() {
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
func (s *suiteUsers) TearDownSuite() {
	config.DB = nil
	s.mock = nil
}

// func (s *suiteUsers) TestLoginUserController() {
// 	var testCases = []struct {
// 		name       string
// 		path       string
// 		expectCode int
// 		method     string
// 		message    string
// 		Body       model.User
// 	}{
// 		{
// 			name:       "login normal",
// 			path:       "/users/login",
// 			expectCode: http.StatusOK,
// 			method:     http.MethodPost,
// 			message:    "user created!",
// 			Body: model.User{
// 				Email:    "email@example.com",
// 				Password: "password",
// 			},
// 		},
// 	}
// 	for _, v := range testCases {
// 		s.T().Run(v.name, func(t *testing.T) {
// 			if v.expectCode == http.StatusOK {
// 				expectedRow := s.mock.NewRows([]string{"id", "name", "email"}).
// 					AddRow(1, "Example", "email@example.com")
// 				s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE (email = ? AND password = ?) AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1")).
// 					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
// 					WillReturnRows(expectedRow)
// 			}
// 			res, _ := json.Marshal(v.Body)
// 			r := httptest.NewRequest(v.method, "/", bytes.NewBuffer(res))
// 			w := httptest.NewRecorder()
// 			e := echo.New()
// 			ctx := e.NewContext(r, w)
// 			ctx.SetPath(v.path)
// 			if s.NoError(LoginUserController(ctx)) {
// 				body := w.Body.Bytes()
// 				type Response struct {
// 					Message string     `json:"message"`
// 					Data    model.User `json:"data"`
// 				}
// 				var user Response
// 				if err := json.Unmarshal(body, &user); err != nil {
// 					s.Error(err, "error unmarshalling")
// 				}
// 				s.Equal(v.expectCode, w.Result().StatusCode)
// 				s.Equal(v.message, user.Message)
// 				s.Equal(v.Body, body)
// 			}
// 		})
// 	}
// }

func (s *suiteUsers) TestGetUsersController() {

	var testCases = []struct {
		name       string
		path       string
		method     string
		expectCode int
		sizeData   int
		message    string
		Body       []model.User
	}{
		{
			name:       "get users normal",
			path:       "/users",
			method:     http.MethodGet,
			expectCode: http.StatusOK,
			sizeData:   2,
			message:    "success",
			Body: []model.User{
				{
					Name:     "User1",
					Email:    "user1@example.com",
					Password: "password1",
				},
				{
					Name:     "User2",
					Email:    "user2@example.com",
					Password: "password2",
				},
			},
		},
	}

	for _, v := range testCases {
		s.T().Run(v.name, func(t *testing.T) {
			expectedResult := sqlmock.NewRows([]string{"name", "email", "password"}).
				AddRow("User1", "user1@example.com", "password1").
				AddRow("User2", "user2@example.com", "password2")
			s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL")).
				WillReturnRows(expectedResult)

			res, _ := json.Marshal(v.Body)

			r := httptest.NewRequest(v.method, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()
			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath(v.path)

			if s.NoError(GetUsersController(ctx)) {
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
				s.Equal(v.Body, user.Data)
			}
		})
	}
}
func (s *suiteUsers) TestGetUsersControllerInvalid() {

	var testCases = []struct {
		name           string
		path           string
		method         string
		expectedResult string
	}{
		{
			name:           "get users database error",
			path:           "/users",
			method:         http.MethodGet,
			expectedResult: echo.NewHTTPError(http.StatusInternalServerError, "Record Not found").Error(),
		},
	}

	for _, v := range testCases {
		s.T().Run(v.name, func(t *testing.T) {

			s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL")).
				WillReturnError(errors.New("Record Not found"))

			r := httptest.NewRequest(v.method, "/", nil)
			w := httptest.NewRecorder()
			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath(v.path)

			err := GetUsersController(ctx).Error()
			s.Equal(v.expectedResult, err)
		})
	}
}

func (s *suiteUsers) TestGetUserController() {

	testCases := []struct {
		name       string
		path       string
		expectCode int
		message    string
		Body       model.User
	}{
		{
			name:       "get user normal",
			path:       "/:id",
			expectCode: http.StatusOK,
			message:    "success",
			Body: model.User{
				Name:     "User",
				Email:    "user@example.com",
				Password: "password",
			},
		},
	}

	for _, v := range testCases {
		s.T().Run(v.name, func(t *testing.T) {

			expectedResult := sqlmock.NewRows([]string{"name", "email", "password"}).
				AddRow("User", "user@example.com", "password")
			s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE id = ? AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1")).
				WithArgs(1).
				WillReturnRows(expectedResult)

			res, _ := json.Marshal(v.Body)

			r := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()
			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath(v.path)
			ctx.SetParamNames("id")
			ctx.SetParamValues("1")

			if s.NoError(GetUserController(ctx)) {
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
				s.Equal(v.Body, user.Data)
			}
		})
	}
}
func (s *suiteUsers) TestGetUserControllerInvalid() {

	var testCases = []struct {
		name           string
		path           string
		method         string
		expectedResult string
	}{
		{
			name:           "get user error",
			path:           "/users/:id",
			method:         http.MethodGet,
			expectedResult: echo.NewHTTPError(http.StatusInternalServerError, "Record Not found").Error(),
		},
	}

	for _, v := range testCases {
		s.T().Run(v.name, func(t *testing.T) {

			s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE id = ? AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1")).
				WillReturnError(errors.New("Record Not found"))

			r := httptest.NewRequest(v.method, "/", nil)
			w := httptest.NewRecorder()
			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath(v.path)
			ctx.SetParamNames("id")
			ctx.SetParamValues("1")

			err := GetUserController(ctx).Error()
			s.Equal(v.expectedResult, err)
		})
	}
}

func (s *suiteUsers) TestCreateUserController() {

	var testCases = []struct {
		name       string
		path       string
		expectCode int
		method     string
		message    string
		Body       model.User
	}{
		{
			name:       "create users normal",
			path:       "/users",
			expectCode: http.StatusCreated,
			method:     http.MethodPost,
			message:    "user created!",
			Body: model.User{
				Name:     "Udin",
				Email:    "udin@test.com",
				Password: "P",
			},
		},
	}

	for _, v := range testCases {
		s.T().Run(v.name, func(t *testing.T) {

			// if v.expectCode == http.StatusCreated {
			s.mock.ExpectBegin()
			s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `users` (`created_at`,`updated_at`,`deleted_at`,`name`,`email`,`password`,`token`) VALUES (?,?,?,?,?,?,?)")).
				WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), nil, "Udin", "udin@test.com", "P", "").
				WillReturnResult(sqlmock.NewResult(1, 1))
			s.mock.ExpectCommit()
			// }

			res, _ := json.Marshal(v.Body)
			r := httptest.NewRequest(v.method, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()
			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath(v.path)
			ctx.Request().Header.Set("Content-Type", "application/json")

			if s.NoError(CreateUserController(ctx)) {
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
func (s *suiteUsers) TestCreateUserControllerInvalid() {

	var testCases = []struct {
		name           string
		path           string
		method         string
		expectedResult string
	}{
		{
			name:           "create user error",
			path:           "/users",
			method:         http.MethodPost,
			expectedResult: echo.NewHTTPError(http.StatusInternalServerError, "Record Not found").Error(),
		},
	}

	for _, v := range testCases {
		s.T().Run(v.name, func(t *testing.T) {

			s.mock.ExpectBegin()
			s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `users` (`created_at`,`updated_at`,`deleted_at`,`name`,`email`,`password`,`token`) VALUES (?,?,?,?,?,?,?)")).
				WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), nil, "", "", "", "").
				WillReturnError(errors.New("Record Not found"))
			s.mock.ExpectRollback()

			r := httptest.NewRequest(v.method, "/", nil)
			w := httptest.NewRecorder()
			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath(v.path)
			ctx.Request().Header.Set("Content-Type", "application/json")

			err := CreateUserController(ctx).Error()
			s.Equal(v.expectedResult, err)
		})
	}
}

func (s *suiteUsers) TestUpdateUserController() {
	testCases := []struct {
		name       string
		path       string
		expectCode int
		method     string
		message    string
		Body       model.User
	}{
		{
			name:       "update user normal",
			path:       "/users/:id",
			expectCode: http.StatusOK,
			method:     http.MethodPut,
			message:    "user updated!",
			Body:       model.User{},
		},
	}

	for _, v := range testCases {
		s.T().Run(v.name, func(t *testing.T) {
			if v.expectCode == http.StatusOK {
				s.mock.ExpectBegin()
				s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `users` SET `updated_at`=? WHERE users.id = ? AND `users`.`deleted_at` IS NULL")).
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

			if s.NoError(UpdateUserController(ctx)) {
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
func (s *suiteUsers) TestUpdateUserControllerInvalid() {

	var testCases = []struct {
		name                  string
		path                  string
		method                string
		expectedResult        string
		expectedStrconvResult string
	}{
		{
			name:           "update user error",
			path:           "/users/:id",
			method:         http.MethodPut,
			expectedResult: echo.NewHTTPError(http.StatusInternalServerError, "Record Not found").Error(),
		},
	}

	for _, v := range testCases {
		s.T().Run(v.name, func(t *testing.T) {

			s.mock.ExpectBegin()
			s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `users` SET `updated_at`=? WHERE users.id = ? AND `users`.`deleted_at` IS NULL")).
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

			err := UpdateUserController(ctx1).Error()
			s.Equal(v.expectedResult, err)
		})
	}
}

func (s *suiteUsers) TestDeleteUserController() {
	testCases := []struct {
		name       string
		path       string
		expectCode int
		method     string
		message    string
	}{
		{
			name:       "delete user normal",
			path:       "/users/:id",
			expectCode: http.StatusOK,
			method:     http.MethodDelete,
			message:    "delete success",
		},
	}

	for _, v := range testCases {
		s.T().Run(v.name, func(t *testing.T) {
			if v.method == http.MethodDelete {
				s.mock.ExpectBegin()
				s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `users` SET `deleted_at`=? WHERE `users`.`id` = ? AND `users`.`deleted_at` IS NULL")).
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

			if s.NoError(DeleteUserController(ctx)) {
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
func (s *suiteUsers) TestDeleteUserControllerInvalid() {

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
			s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `users` SET `deleted_at`=? WHERE `users`.`id` = ? AND `users`.`deleted_at` IS NULL")).
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

			err := DeleteUserController(ctx).Error()
			s.Equal(v.expectedResult, err)
		})
	}
}
