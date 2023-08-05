package service

import (
	"golang-gin/pkg/data/request"
	"golang-gin/pkg/data/response"
)

type GiftService interface {
	Create(gift request.CreateGiftRequest) response.GiftResponse
	Update(gift request.UpdateGiftRequest)
	Delete(giftId int)
	FindById(giftId int) request.UpdateGiftRequest
	FindAll() []response.GiftResponse
}
