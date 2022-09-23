package main

import (
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

// get all users
func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	// your solution here
	id, _ := strconv.Atoi(c.Param("id"))
	var result User
	for i := range users {
		user := users[i]
		if id == user.Id {
			result = user
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get user by id: " + strconv.Itoa(id),
		"user":     result,
	})
}

func DeleteUserController(c echo.Context) error {
	// your solution here
	id, _ := strconv.Atoi(c.Param("id"))
	tempUsers := users
	ok := false
	for i := range users {
		if id == tempUsers[i].Id {
			for j := i; j != len(tempUsers); j++ {
				tempUsers[j].Name = tempUsers[j+1].Name
				tempUsers[j].Email = tempUsers[j+1].Email
				tempUsers[j].Password = tempUsers[j+1].Password
			}
			ok = true
			break
		}
	}

	if !ok {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"messages": "Data with id: " + strconv.Itoa(id) + " not found",
		})
	} else {
		if len(users) > 1 {
			users = tempUsers[:len(tempUsers)-1]
		} else {
			users = tempUsers
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"messages": "Data with id: " + strconv.Itoa(id) + " deleted",
		})
	}
}

func UpdateUserController(c echo.Context) error {
	// your solution here
	id, _ := strconv.Atoi(c.Param("id"))
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

// create new user
func CreateUserController(c echo.Context) error {
	user := User{}
	c.Bind(&user)

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
	e.GET("/users", GetUsersController)
	e.POST("/users", CreateUserController)

	//get user by id
	e.GET("/users/:id", GetUserController)

	//delete user by id
	e.DELETE("/users/:id", DeleteUserController)

	//update user (PUT) information by id
	e.PUT("/users/:id", UpdateUserController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
