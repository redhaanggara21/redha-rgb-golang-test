package request

import "time"

type UpdateGiftRequest struct {
	Id          int    `validate:"required"`
	Name        string `validate:"required,min=1,max=200" json:"name"`
	Description string `validate:"required,min=1,max=200,description" json:"description"`
	Image       string `json:"image"`
	Stock       int    `json:"stock"`
}

type Files struct {
	Id       int64
	User_Id  int64
	FileName string
	Type     string
	Name     string
	Link     string
	Created  time.Time
	Updated  time.Time
}
