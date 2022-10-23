package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `json:"email" form:"email"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Blogs    []Blog
}
