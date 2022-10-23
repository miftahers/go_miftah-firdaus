package services

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

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserSuites struct {
	suite.Suite
	mock    sqlmock.Sqlmock
	service UserService
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
	_, uServices, _ := NewServices(repo)
	s.service = uServices
	s.mock = mock
}

func (s *UserSuites) TearDownSuite() {
	s.mock = nil
}

func (s *UserSuites) TestSignup() {
	test := []struct {
		name         string
		expectResult error
		body         models.User
		method       string
		path         string
	}{
		{
			name:         "signup service normal",
			expectResult: nil,
			body: models.User{
				Email:    "user@test.com",
				Username: "user1",
				Password: "123123123",
			},
			method: http.MethodPost,
			path:   "/blogs/category",
		},
	}

	for _, v := range test {
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

			err := s.service.SignUp(c)

			s.Equal(v.expectResult, err)
		})
	}
}

func (s *UserSuites) TestSignupError() {
	test := []struct {
		name         string
		expectResult error
		body         models.User
		method       string
		path         string
	}{
		{
			name:         "signup service normal",
			expectResult: errors.New("Database error"),
			body: models.User{
				Email:    "user@test.com",
				Username: "user1",
				Password: "123123123",
			},
			method: http.MethodPost,
			path:   "/blogs/category",
		},
	}

	for _, v := range test {
		s.T().Run(v.name, func(t *testing.T) {

			s.mock.ExpectBegin()
			s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `users` (`created_at`,`updated_at`,`deleted_at`,`email`,`username`,`password`) VALUES (?,?,?,?,?,?)")).
				WithArgs(AnyTime{}, AnyTime{}, nil, "user@test.com", "user1", "123123123").
				WillReturnError(errors.New("Database error"))
			s.mock.ExpectRollback()

			res, _ := json.Marshal(v.body)

			r := httptest.NewRequest(v.method, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()

			c := echo.New().NewContext(r, w)
			c.SetPath(v.path)
			c.Request().Header.Set("Content-Type", "application/json")

			err := s.service.SignUp(c)

			s.Equal(v.expectResult, err)
		})
	}
}

func (s *UserSuites) TestSigIn() {
	test := []struct {
		name        string
		expectError error
		expectBody  dto.LoginUser
		body        models.User
		method      string
		path        string
	}{
		{
			name:        "signup service normal",
			expectError: nil,
			expectBody: dto.LoginUser{
				Username: "user1",
			},
			body: models.User{
				Email:    "user@test.com",
				Username: "user1",
				Password: "123123123",
			},
			method: http.MethodPost,
			path:   "/blogs/category",
		},
	}

	for _, v := range test {
		s.T().Run(v.name, func(t *testing.T) {

			row := s.mock.NewRows([]string{"user_id", "username"}).AddRow(1, "user1")
			s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE (username = ? AND password = ?) AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1")).
				WithArgs("user1", "123123123").
				WillReturnRows(row)

			res, _ := json.Marshal(v.body)

			r := httptest.NewRequest(v.method, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()

			c := echo.New().NewContext(r, w)
			c.SetPath(v.path)
			c.Request().Header.Set("Content-Type", "application/json")

			result, err := s.service.Login(c)
			if err != nil {
				s.Error(err)
			}

			s.Equal(v.expectBody.Username, result.Username)
		})
	}
}

func (s *UserSuites) TestSigInError() {
	test := []struct {
		name        string
		expectError error
		expectBody  dto.LoginUser
		body        models.User
		method      string
		path        string
	}{
		{
			name:        "signup service normal",
			expectError: errors.New("record not found"),
			expectBody: dto.LoginUser{
				Username: "user1",
			},
			body: models.User{
				Email:    "user@test.com",
				Username: "user1",
				Password: "123123123",
			},
			method: http.MethodPost,
			path:   "/blogs/category",
		},
	}

	for _, v := range test {
		s.T().Run(v.name, func(t *testing.T) {

			s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE (username = ? AND password = ?) AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1")).
				WithArgs("user1", "123123123").
				WillReturnError(errors.New("record not found"))

			res, _ := json.Marshal(v.body)

			r := httptest.NewRequest(v.method, "/", bytes.NewBuffer(res))
			w := httptest.NewRecorder()

			c := echo.New().NewContext(r, w)
			c.SetPath(v.path)
			c.Request().Header.Set("Content-Type", "application/json")

			_, err := s.service.Login(c)

			s.Equal(v.expectError, err)
		})
	}
}
