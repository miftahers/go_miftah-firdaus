package model

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title        string `json:"title" form:"title"`
	Category     string `json:"category" form:"category"`
	Release_Year string `json:"release_year" form:"release_year"`
	Writter      string `json:"writter" form:"writter"`
}
