package service

import (
	"golang-crud-gin/data/request"
	"golang-crud-gin/data/response"
	"golang-crud-gin/helper"
	"golang-crud-gin/model"
	"golang-crud-gin/repository"

	"github.com/go-playground/validator/v10"
)

type ReedemServiceImpl struct {
	ReedemRepository repository.ReedemRepository
	Validate         *validator.Validate
}

func NewReedemServiceImpl(reedemRepository repository.ReedemRepository, validate *validator.Validate) ReedemService {
	return &ReedemServiceImpl{
		ReedemRepository: reedemRepository,
		Validate:         validate,
	}
}

// Create implements ReedemService.
func (t *ReedemServiceImpl) Create(reedem request.CreateReedemRequest) response.ReedemResponse {
	err := t.Validate.Struct(reedem)
	helper.ErrorPanic(err)
	reedemModel := model.Reedem{
		UserID: reedem.UserID,
		GiftID: reedem.GiftID,
		Amount: reedem.Amount,
	}
	data, err := t.ReedemRepository.Save(reedemModel)

	helper.ErrorPanic(err)

	reedemResponse := response.ReedemResponse{
		Id: int(data.ID),
	}

	return reedemResponse
}

// Delete implements ReedemService.
func (t *ReedemServiceImpl) Delete(reedemId int) {
	t.ReedemRepository.Delete(reedemId)
}

// FindAll implements ReedemService.
func (t *ReedemServiceImpl) FindAll() []response.ReedemResponse {
	result := t.ReedemRepository.FindAll()

	var reedems []response.ReedemResponse
	for _, value := range result {
		reedem := response.ReedemResponse{
			Id:     int(value.ID),
			GiftID: value.GiftID,
			UserID: value.UserID,
			Amount: value.Amount,
		}
		reedems = append(reedems, reedem)
	}

	return reedems
}

// FindById implements ReedemService.
func (t *ReedemServiceImpl) FindById(reedemId int) response.ReedemResponse {
	reedemData, err := t.ReedemRepository.FindById(reedemId)
	helper.ErrorPanic(err)

	reedemResponse := response.ReedemResponse{
		Id:     int(reedemData.ID),
		UserID: reedemData.UserID,
		GiftID: reedemData.GiftID,
		Amount: reedemData.Amount,
	}
	return reedemResponse
}

// Update implements ReedemService.
func (t *ReedemServiceImpl) Update(reedem request.UpdateReedemRequest) {
	reedemData, err := t.ReedemRepository.FindById(reedem.Id)
	helper.ErrorPanic(err)
	reedemData.GiftID = reedem.GiftID
	reedemData.UserID = reedem.UserID
	reedemData.Amount = reedem.Amount
	t.ReedemRepository.Update(reedemData)
}
