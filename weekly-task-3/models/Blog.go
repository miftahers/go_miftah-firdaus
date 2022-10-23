package models

import (
	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model
	UUID       string `gorm:"primaryKey"`
	Title      string `json:"title" form:"title"`
	Content    string `json:"content" form:"content"`
	UserID     uint   `json:"user_id" form:"user_id"`
	CategoryID uint   `json:"category_id" form:"category_id"`
	Category   Category
}
