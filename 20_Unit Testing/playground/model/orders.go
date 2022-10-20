package model

import "gorm.io/gorm"

type Orders struct {
	gorm.Model
	OrderID string `json:"order_id" form:"order_id"`
	QTY     int    `json:"qty" form:"qty"`
}
