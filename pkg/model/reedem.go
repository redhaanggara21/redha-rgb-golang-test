package model

import "gorm.io/gorm"

type Reedem struct {
	gorm.Model
	UserID int `json:"user_id"`
	GiftID int `json:"gift_id"`
	Amount int `json:"amount"`
	Test   int `json:"test"`
}
