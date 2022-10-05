package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"restful-api-testing/config"
	"restful-api-testing/model"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func InitEchoTestAPI() *echo.Echo {
	config.InitDBTest()
	e := echo.New()
	return e
}

func InsertDataUserForGetUsers() error {
	user := model.User{
		Name:     "Naruto Uzumaki",
		Password: "Hinata123",
		Email:    "NarutoAndHinata@forever.com",
	}
	if err := config.DB.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func TestGetUserController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
		sizeData   int
	}{
		{
			name:       "get user normal",
			path:       "/users",
			expectCode: http.StatusOK,
			sizeData:   1,
		},
	}

	e := InitEchoTestAPI()
	InsertDataUserForGetUsers()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, GetUserController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()

			//open file
			//convert struct
			var user UserResponse
			if err := json.Unmarshal([]byte(body), &user); err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, testCase.sizeData, len(user.Data))
		}
	}
}
