package service

import (
	"golang-gin/pkg/data/request"
	"golang-gin/pkg/data/response"
)

type RateService interface {
	Create(rate request.CreateRateRequest) response.RateResponse
	Update(rate request.UpdateRateRequest)
	Delete(rateId int)
	FindById(rateId int) response.RateResponse
	FindAll() []response.RateResponse
}
