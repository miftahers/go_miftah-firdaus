# Middleware

- What is middleware?
- Implementation middleware
- Echo middleware
- type echo middleware
- implementation
  - Log Middleware
  - Auth Middleware
  - JWT Middleware

## What is middleware?

Middleware adalah sebuah entity yang dipasangkan pada sebuah proses server. Ketika client membuat sebuah request dan server memberikan respon, akan diberikan sebuah layer atau middleware yang berisi fungsi-fungsi khusus. Fungsi-fungsi khusus tersebut akan membantu komunikasi data antara client dan server.

```
HTTP Request -> Middleware -> API Server

HTTP Response <- Middleware <- API Server
```

Example Third Party Middleware
- Negroni
- Echo
- Interpose
- Alice
- Or make ur own...

## Echo Middleware

### Contents

- Basic Auth
- Body Dump
- Body Limit
- CORS
- CSRF
- Casbin Auth
- Gzip
- JWT
- Key Auth
- Logger
- Method Override
- Proxy
- Recover
- Redirect
- Request ID
- Rewrite
- Secure
- Session
- Static
- Trailing Slash

### Dokumentasi Echo Middleware 

https://echo.labstack.com/middleware/

## Type Echo Middleware
### Echo#Pre()
Echo#Pre() -> Executed before router processes the request

- HTTPSRedirect
- HTTPSWWWRedirect
- WWWRedirect
- AddTrailingSlash
- RemoveTrailingSlash
- MethodOverride
- Rewrite

### Echo#Use()
Echo#Use() -> Executed after router processes the requesst and has full access to echo.Context API (Biasanya ada di controller)

- BodyLimit
- Logger
- Gzip
- Recover
- BasicAuth
- JWTAuth
- Secure
- CORS
- Static

#### HTTPS REDIRECT

http request to https OR mengubah http menjadi https

contoh:

`http://labstack.com`

akan di redirect ke

`https://labstack.com`

```go
e := echo.New()
e.Pre(middleware.HTTPSRedirect())
```

## Implementation

### Log Middleware

- Mencatat informasi apa saja yang terjadi di server kita, terutama di dalam request (Logs the information HTTP request)
- Sebagai rekam jejak (As a footprint)
- Sumber data untuk analisik(Datasource for analytics)

```go
// middleware folder create log_middleware.go and write this
package middleware

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func LogMiddleware(e *echo.Echo) {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status},latency_human=${latency_human}, remote_ip=${remote_ip}\n",
	}))
}

// then in route file / routes.go call the LogMiddleware Func with echo param. Write this func as example

func New() *echo.Echo {
	e := echo.New()

	// routing
	e.GET("/users", controller.GetUserController)
  //Call LogMiddleware here
	middleware.LogMiddleware(e)
	e.POST("/users", controller.CreateUserController)
	return e
}
```

### Auth Middleware

Why using authentication?
- User Identification
- Secure Data and Information

#### Basic Authentication

adalah proses authentikasi pada REST API dengan mengirimkan data username dan password melalui request headers.

Format:
- Information Encoded Format:
  - 'Authorization: Basic ' + base64encode('username:password')
- Header Generate:
  - Authorization: Basic dXN1cm5hbWU6cGFzc4dvcmQ=


```go
// di folder middleware create auth_middleware.go and write this
package middleware

func BasicAuthDB(username, password string, c echo.Context) (bool, error) {
	// Be careful to use constant time comparison to prevent timing attacks
	var user model.User
	err := config.DB.Where("email= ? AND password= ?", username, password).First(&user).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

// kemudian di bagian routes buat sebuah group endpoint sebagai contoh dan tulis sebagai berikut
	eAuthBasic := e.Group("/auth")
	// middleware.BasicAuth memanggil fungsi di package echo.Middleware sedangkan m.BasicAuthDB memanggil fungsi di package middleware lokal
	eAuthBasic.Use(middleware.BasicAuth(m.BasicAuthDB)) 
	eAuthBasic.GET("/users", controller.GetUserController)
```


### JWT Middleware

Adalah sebuah autentifikasi seperti basic auth, tetapi algoritma nya berbeda. Jika basic auth menggunakan username dan password untuk dikirim ke server dan server mengecek ke database. Masalah yang timbul dari basic auth adalah ketika ada 1 juta request ke server maka server kita akan melakukan 1 juta pengecekan ke database dan akan membebani server kita. Untuk menanggulangi hal tersebut kita bisa menggunakan JWT middleware dimana validasi hanya dilakukan di server saja.

JWT memiliki beberapa token yang dipisahkan oleh titik (".").
- Token pertama (HEADER) adalah hasil encoding64 sebuah json informasi terkait algoritma yang akan digunakan.
- Token yang kedua (PAYLOAD) adalah hasil encoding64 dari sebuah json yang bisa lebih customize. Contoh kita bisa mengirimkan data sub, name, int, dll. Selain itu jika kita ingin token ini expired dalam kurun waktu tertentu kita bisa menambahkan parameter "exp"/"expired" pada json.
- Token yang ketiga adalah hasil encoding64 dari HEADER + PAYLOAD + kunci dari kita (contoh: secret). Lalu kita akan melakukan sebuah algoritma untuk memvalidasi token.

Format:
- Information Encoded Format:
  - 'Authorization: Bearer + Token
- Header Generate:
  - Authorization: Bearer eyJhbGci0iJIUzI1NiIsInR5cCI...

Example in playground folder.