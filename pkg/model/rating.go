package model

import "gorm.io/gorm"

type Rating struct {
	gorm.Model
	UserID int    `json:"user_id"`
	GiftID int    `json:"gift_id"`
	Rate   int    `json:"rate"`
	Test   string `json:"test"`
}
