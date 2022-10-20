package model

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Name           string `json:"name" form:"name"`
	Email          string `json:"email" form:"email"`
	ProfilePicture ProfilePicture
}
