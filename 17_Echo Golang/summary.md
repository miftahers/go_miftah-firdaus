# Echo Golang

Echo adalah web framework golang yang mempunyai performa tinggi, extensible (bisa dipasang-pasangkan), dan minimalist. 

Kelebihan dari echo framework:
1. Optimized Router
2. Middleware
3. Data Rendering
4. Scalable
5. Data Binding

Echo it`s a minimalist framework
- no database driver or ORM
- No folder scaffold provided, define your own structure
- template engine: https://echo.labstack.com/guide/templates

Minimalist but EXTENSIBLE
- GORM for ORM (https://gorm.io)
- Go JWT Extended for Authentication (https://github.com/dgrijalva/jwt-go)
- etc...

Tutorial pake echo golang:
1. buat folder `rest`
2. buat go mod `rest` dengan ketik `go mod init rest` di terminal
3. inisialisasi echo dengan mengetik `go get -u github.com/labstack/echo/...`
4. buat file main.go

## Routing dan controller

Contoh router -> `https://localhost:8080/user`
Ketika request method nya adalah GET maka router akan mengarahkan ke controller dengan fungsi getUser

### Rendering data
```go
// echo.String -> rendering data ke String

package main

import (
  "net/http"
  "github.com/labstack/echo"
)
func main() {
  e := echo.New()

  // routing
  e.GET("/user", HelloController)

  e.Start(":8080")
}

// Response "Hello World"
func HelloController(e echo.Context) error { // echo.Context berfungsi untuk menerima data dari client, selain itu echo.Context juga berfungsi memberikan respon ke client.
  return e.String(http.StatusOK, "Hello World")
}
```

```go
// echo.JSON -> rendering data ke JSON

package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type User struct {
	Email string
	Name  string
}

func main() {
	e := echo.New()

	// routing
	e.GET("/User", UserController)

	e.Start(":8080")
}

func UserController(e echo.Context) error {
	user := User{
		Email: "miftahers@gmail.com",
		Name:  "Alta",
	}

	return e.JSON(http.StatusOK, user)
}
```

### Retrieving data
Ada beberapa cara client mendapat data, contohnya:
#### url params

```go
package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type User struct {
	Id    int
	Age   int
	Email string
	Name  string
}

func main() {
	e := echo.New()

	// routing
	e.GET("/user/:id/:age", UserController) // titik dua di depan id & age membuat kita bisa menangkap nilai id & age yang dimaksud client

	e.Start(":8080")
}

func UserController(e echo.Context) error {

	id, _ := strconv.Atoi(e.Param("id"))
	age, _ := strconv.Atoi(e.Param("age"))

	user := User{
		Id:    id,
		Age:   age,
		Email: "miftahers@gmail.com",
		Name:  "Alta",
	}

	return e.JSON(http.StatusOK, user)
}
```
#### query params

```go
package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type User struct {
	Id    int
	Age   int
	Email string
	Name  string
}

func main() {
	e := echo.New()

	// routing
	e.GET("/user/:id/:age", UserController) 

	e.Start(":8080")
}

// controller
func UserController(e echo.Context) error {

	id, _ := strconv.Atoi(e.Param("id"))
	age, _ := strconv.Atoi(e.Param("age"))

	search := e.QueryParam("search")
	sort := e.QueryParam("urutan")

	user := User{
		Id:    id,
		Age:   age,
		Email: "miftahers@gmail.com",
		Name:  "Alta",
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"user":   user,
		"search": search,
		"sort":   sort,
	})
}

```

#### Handle POST using echo.formValue()

```go
package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type User struct {
	Id    int
	Age   int
	Email string
	Name  string
}

func main() {
	e := echo.New()

	// routing
	e.POST("/user", RegisterController)
	e.Start(":8080")
}

// controller
func RegisterController(c echo.Context) error {
	email := c.FormValue("email")
	name := c.FormValue("name")

	return c.JSON(http.StatusOK, map[string]interface{}{
		"email": email,
		"name":  name,
	})
}
```

#### Handle POST using echo.Bind() for JSON/FORM
```go

package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type User struct {
	Id    int    `json:"id" form:"id"`
	Age   int    `json:"age" form:"age"`
	Email string `json:"email" form:"email"`
	Name  string `json:"name" form:"name"`
}

func main() {
	e := echo.New()

	// routing
	e.POST("/user", RegisterController)
	e.Start(":8080")
}

// controller
func RegisterController(c echo.Context) error {
	// buat variable struct User
	user := User{}

	// bind data yang masuk ke struct User
	c.Bind(&user)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"user": user,
	})
}

```

