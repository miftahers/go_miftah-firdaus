package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
	"weekly-task-3/dto"
	"weekly-task-3/models"
	"weekly-task-3/repositories"
	"weekly-task-3/services"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserSuites struct {
	suite.Suite
	mock    sqlmock.Sqlmock
	handler UserHandler
}

func TestSuiteUser(t *testing.T) {
	suite.Run(t, new(UserSuites))
}

func (s *UserSuites) SetupSuite() {
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
	_, userServices, _ := services.NewServices(repo)
	handler := UserHandler{
		UserService: userServices,
	}

	s.handler = handler
	s.mock = mock
}

func (s *UserSuites) TearDownSuite() {
	s.mock = nil
}

func (s *UserSuites) TestSignUp() {
	testCase := []struct {
		name       string
		path       string
		method     string
		expectCode int
		message    string
		body       models.User
	}{
		{
			name:       "sign up normal",
			path:       "/signup",
			method:     http.MethodPost,
			expectCode: http.StatusCreated,
			message:    "created",
			body: models.User{
				Email:    "user@test.com",
				Username: "user1",
				Password: "123123123",
			},
		},
	}

	for _, v := range testCase {
		s.T().Run(v.name, func(t *testing.T) {
			s.mock.ExpectBegin()
			s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `users` (`created_at`,`updated_at`,`deleted_at`,`email`,`username`,`password`) VALUES (?,?,?,?,?,?)")).
				WithArgs(AnyTime{}, AnyTime{}, nil, "user@test.com", "user1", "123123123").
				WillReturnResult(sqlmock.NewResult(1, 1))
			s.mock.ExpectCommit()

			res, _ := json.Marshal(v.body)

			r := httptest.NewRequest(v.method, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()
			c := echo.New().NewContext(r, w)
			c.SetPath(v.path)
			c.Request().Header.Set("Content-Type", "application/json")

			if s.NoError(s.handler.SignUp(c)) {
				body := w.Body.Bytes()

				type response struct {
					Message string `json:"message"`
				}
				var resp response
				err := json.Unmarshal(body, &resp)
				if err != nil {
					s.Error(err, "error unmarshalling")
				}

				s.Equal(v.expectCode, w.Result().StatusCode)
				s.Equal(v.message, resp.Message)
			}
		})
	}
}

func (s *UserSuites) TestSignupError() {
	testCase := []struct {
		name        string
		path        string
		method      string
		expectError string
		body        models.User
	}{
		{
			name:        "Sign up - Error",
			path:        "/signup",
			method:      http.MethodPost,
			expectError: echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error").Error(),
			body: models.User{
				Email:    "user@test.com",
				Username: "user1",
				Password: "123123123",
			},
		},
	}

	for _, v := range testCase {
		s.T().Run(v.name, func(t *testing.T) {
			s.mock.ExpectBegin()
			s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `users` (`created_at`,`updated_at`,`deleted_at`,`email`,`username`,`password`) VALUES (?,?,?,?,?,?)")).
				WithArgs(AnyTime{}, AnyTime{}, nil, "user@test.com", "user1", "123123123").
				WillReturnError(errors.New("Internal Server Error"))
			s.mock.ExpectRollback()

			res, _ := json.Marshal(v.body)
			r := httptest.NewRequest(v.method, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()
			c := echo.New().NewContext(r, w)
			c.SetPath(v.path)
			c.Request().Header.Set("Content-Type", "application/json")

			err := s.handler.SignUp(c)

			s.Equal(v.expectError, err.Error())
		})
	}
}

func (s *UserSuites) TestSignin() {

	testCase := []struct {
		name       string
		path       string
		method     string
		body       models.User
		expectBody dto.LoginUser
	}{
		{
			name:   "Sign in normal",
			path:   "/login",
			method: http.MethodPost,
			body: models.User{
				Username: "user1",
				Password: "123123123",
			},
			expectBody: dto.LoginUser{
				Username: "user1",
			},
		},
	}

	for _, v := range testCase {
		s.T().Run(v.name, func(t *testing.T) {
			expectRows := s.mock.NewRows([]string{"email", "password"}).AddRow("user1", "123123123")
			s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE (username = ? AND password = ?) AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1")).
				WithArgs("user1", "123123123").
				WillReturnRows(expectRows)

			res, _ := json.Marshal(v.body)
			r := httptest.NewRequest(v.method, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()
			c := echo.New().NewContext(r, w)
			c.SetPath(v.path)
			c.Request().Header.Set("Content-Type", "application/json")

			if s.NoError(s.handler.Login(c)) {
				body := w.Body.Bytes()

				type reponse struct {
					Message string        `json:"message"`
					Data    dto.LoginUser `json:"data"`
				}
				var resp reponse
				err := json.Unmarshal(body, &resp)
				if err != nil {
					s.Error(err, "error unmarshalling")
				}

				s.Equal(v.expectBody.Username, resp.Data.Username)
			}
		})
	}
}

func (s *UserSuites) TestSigninError() {

	testCase := []struct {
		name        string
		path        string
		method      string
		body        models.User
		expectBody  dto.LoginUser
		expectError string
	}{
		{
			name:   "Sign in Error",
			path:   "/login",
			method: http.MethodPost,
			body: models.User{
				Username: "user1",
				Password: "123123123",
			},
			expectError: echo.NewHTTPError(http.StatusInternalServerError, "Record Not Found").Error(),
		},
	}

	for _, v := range testCase {
		s.T().Run(v.name, func(t *testing.T) {
			s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE (username = ? AND password = ?) AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1")).
				WillReturnError(errors.New("Record Not Found"))

			res, _ := json.Marshal(v.body)
			r := httptest.NewRequest(v.method, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()
			c := echo.New().NewContext(r, w)
			c.SetPath(v.path)
			c.Request().Header.Set("Content-Type", "application/json")

			err := s.handler.Login(c).Error()
			s.Equal(v.expectError, err)
		})
	}
}
