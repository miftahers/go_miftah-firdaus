package models

import (
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	UUID        string `gorm:"primaryKey" son:"uuid" form:"uuid"`
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"dercription"`
	Stock       int    `json:"stock" form:"stock"`
	Price       int    `json:"price" form:"price"`
	CategoryID  uint   `json:"category_id" form:"category_id"`
}
