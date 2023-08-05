package service

import (
	"golang-crud-gin/data/request"
	"golang-crud-gin/data/response"
)

type RateService interface {
	Create(rate request.CreateRateRequest) response.RateResponse
	Update(rate request.UpdateRateRequest)
	Delete(rateId int)
	FindById(rateId int) response.RateResponse
	FindAll() []response.RateResponse
}
