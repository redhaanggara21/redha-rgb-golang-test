package model

import "gorm.io/gorm"

type Gift struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Stock       int    `json:"stock"`
}
