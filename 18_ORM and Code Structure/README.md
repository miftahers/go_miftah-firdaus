# ORM and Code Structure

Outline:
- What is Object Relational Mapping (ORM)?
- Advantages ORM
- Disadvantages ORM
- What is GORM?
- Database Connection
- Database migration
- CRUD RESTful
- What is MVC?
- Why need structure?
- Structuring project


## ORM

Object Realational Mapping membantu ketika melakukan konversi data, dari data yang berasal dari database ke bentuk object (dalam golang berbentuk struct). 

DATABASE `<->` ORM `<->` Golang Struct

Pengertian ORM di computer science adalah teknik pemrograman untuk mengonversi data antara sistem tipe yang tidak kompatibel menggunakan bahasa pemrograman berorientasi objek.

## ORM Advantages

- Lebih sedikit query yang berulang
- Otomatis mengubah data dari database menjadi object(struct) yang siap digunakan
- Lebih mudah untuk melakukan validasi data sebelum di simpan ke database
- Memiliki fitur cache query, menyimpan sementara data yang baru saja digunakan, memungkinkan proses data lebih cepat.

## ORM Disadvantages

- Menambahkan lapisan dalam kode dan cost proses overhead
- Dalam beberapa kasus, ORM akan menambahkan data relasi yang tidak terlalu penting
- query yang complex akan ditulis lebih sulit jika menggunakan ORM (contoh: >10 table joins)
- Ada beberapa ORM yang tidak mendukung beberapa query database.

## What is GORM?

GORM adalah ORM yang ada dalam bahasa pemrograman Golang

## Database connection

Connection for MySQL using GORM:
```
<username>:<password>@/<db_name>?charset=utf8&parseTime=True&loc=
```

## Database migration

Database migration itu penting buat industri karena ketika ada update fitur atau versi API, Database migration bisa melakukan tracking siapa yang menambahkan fitur/tambahan pada database. Selain itu, database migration bisa rollback versi database jika terdapat error/kendala di versi baru.

### Why DB Migration?

- Lebih simpel dalam update database
- Lebih simpel dalam rollback database
- tracking perubahan pada database structure
- histori database structure akan tercatat pada kode
- selalu kompatibel/cocok dengan perubahan/update aplikasi


## Installasi GORM

import package:
```sh
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite
```

## Praktikum GORM -> installasi GORM dan AutoMigrate
```go
/*
import package:

go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite

*/
package main

import (
	"net/http"

	"github.com/labstack/echo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID    int    `json:"id" form:"id"`
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
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

var DB *gorm.DB

  /*

  template connect to mysql
  
  dsn := "<db_username>:<db_password>@tcp(127.0.0.1:3306)/<db_name>?charset=utf8mb4&parseTime=True&loc=Local"
  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
  }

  refer -> https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
  */
  
func initDB() {

	dsn := "root:@tcp(localhost:3306)/playground?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	// init migrate di init DB untuk membuat table baru
	initMigrate()
}

func initMigrate() {
	DB.AutoMigrate(&User{})
}

func main() {
	initDB()

	e := echo.New()

	// routing
	e.POST("/user", RegisterController)
	e.Start(":8080")
}

```

## Praktikum GORM -> membuat getUserController & createUserController with MySQL DB

```go
package main

import (
	"net/http"

	"github.com/labstack/echo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID    int    `json:"id" form:"id"`
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
}

var DB *gorm.DB

func initDB() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:@tcp(localhost:3306)/playground?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	// init migrate di init DB untuk membuat table baru
	initMigrate()
}

func initMigrate() {
	DB.AutoMigrate(&User{})
}

func main() {
	initDB()

	e := echo.New()

	// routing
	e.GET("/users", GetUserController)
	e.POST("/users", CreateUserController)
	e.Start(":8080")
}

func GetUserController(c echo.Context) error {
	var users []User

	// GORM akan mengecek tabel users dan menyimpan data-datanya ke variabel users
	err := DB.Find(&users).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    users,
	})
}

// controller
func CreateUserController(c echo.Context) error {
	// buat variable struct User
	user := User{}
	// bind data yang masuk ke struct User
	c.Bind(&user)
	//Insert data to DB
	err := DB.Save(&user).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create!",
		"data":    user,
	})
}
```

## Praktikum GORM -> Refactoring Project ke Bentuk yang lebih mudah dilihat dan dibaca (Model View Controller / MVC)
```
Struktur folder:

project >> <main.go, go.mod, go.sum>
- - - Config >> <config.go> (set -> package config | isi nya configurasi DB)
- - - Controller >> <controller.go> (set -> package controller | isinya controller/handler)
- - - Model >> <user.go, dan lainnya> (set -> package model | isinya struct object termasuk user.go berisi struct User)
