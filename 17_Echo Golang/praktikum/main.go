package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User

// GetUsersController -> get all users
func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

// GetUserController -> get user by id
func GetUserController(c echo.Context) error {
	// your solution here
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	var index int
	for i := range users {
		user := users[i]
		if id == user.Id {
			index = i
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get user by id: " + strconv.Itoa(id),
		"user":     users[index],
	})
}

func DeleteUserController(c echo.Context) error {
	// your solution here
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	ok := false
	for i := range users {
		if id == users[i].Id {
			users = append(users[:i], users[i+1:]...)
			ok = true
			break
		}
	}

	if !ok {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"messages": "Data with id: " + strconv.Itoa(id) + " not found",
		})
	} else {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"messages": "Data with id: " + strconv.Itoa(id) + " deleted",
		})
	}
}

func UpdateUserController(c echo.Context) error {
	// your solution here
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")

	var index int
	for i := range users {
		if id == users[i].Id {
			index = i
			users[i].Name = name
			users[i].Email = email
			users[i].Password = password
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success update user by id: " + strconv.Itoa(id),
		"user":     users[index],
	})
}

// CreateUserController -> create new user
func CreateUserController(c echo.Context) error {
	user := User{}
	err := c.Bind(&user)
	if err != nil {
		fmt.Println("ini error")
		return c.String(404, "error")
	}

	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}
	users = append(users, user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})
}

func main() {
	e := echo.New()
	//routing with query parameter
	v1 := e.Group("/v1")
	v1.GET("/users", GetUsersController)
	v1.POST("/users", CreateUserController)

	//get user by id
	v1.GET("/users/:id", GetUserController)

	//delete user by id
	v1.DELETE("/users/:id", DeleteUserController)

	//update user (PUT) information by id
	v1.PUT("/users/:id", UpdateUserController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
