package service

import (
	"golang-crud-gin/data/request"
	"golang-crud-gin/data/response"
)

type ReedemService interface {
	Create(reedem request.CreateReedemRequest) response.ReedemResponse
	Update(reedem request.UpdateReedemRequest)
	Delete(reedemId int)
	FindById(reedemId int) response.ReedemResponse
	FindAll() []response.ReedemResponse
}
