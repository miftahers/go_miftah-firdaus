package model

import (
	"gorm.io/gorm"
)

type ProfilePicture struct {
	gorm.Model
	Path    string `binding:"required" json:"path" form:"path"`
	UsersID uint   `json:"userid" form:"userid"`
}
