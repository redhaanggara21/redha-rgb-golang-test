package service

import (
	"golang-crud-gin/data/request"
	"golang-crud-gin/data/response"
)

type GiftService interface {
	Create(gift request.CreateGiftRequest) response.GiftResponse
	Update(gift request.UpdateGiftRequest)
	Delete(giftId int)
	FindById(giftId int) request.UpdateGiftRequest
	FindAll() []response.GiftResponse
}
