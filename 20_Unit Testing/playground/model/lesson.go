package model

import "gorm.io/gorm"

type Lessons struct {
	gorm.Model
	Title string `json:"title" form:"title"`
	Users []Users
}
