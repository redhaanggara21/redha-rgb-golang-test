package service

import (
	"golang-gin/pkg/data/request"
	"golang-gin/pkg/data/response"
)

type ReedemService interface {
	Create(reedem request.CreateReedemRequest) response.ReedemResponse
	Update(reedem request.UpdateReedemRequest)
	Delete(reedemId int)
	FindById(reedemId int) response.ReedemResponse
	FindAll() []response.ReedemResponse
}
