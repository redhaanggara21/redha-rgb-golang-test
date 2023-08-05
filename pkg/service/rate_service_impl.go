package service

import (
	"golang-gin/pkg/data/request"
	"golang-gin/pkg/data/response"
	"golang-gin/pkg/helper"
	"golang-gin/pkg/model"
	"golang-gin/pkg/repository"

	"github.com/go-playground/validator/v10"
)

type RateServiceImpl struct {
	RateRepository repository.RateRepository
	Validate       *validator.Validate
}

func NewRateServiceImpl(rateRepository repository.RateRepository, validate *validator.Validate) RateService {
	return &RateServiceImpl{
		RateRepository: rateRepository,
		Validate:       validate,
	}
}

// Create implements RateService.
func (t *RateServiceImpl) Create(rate request.CreateRateRequest) response.RateResponse {
	err := t.Validate.Struct(rate)
	helper.ErrorPanic(err)
	ratemModel := model.Rating{
		Rate:   int(rate.Rated),
		UserID: rate.UserID,
		GiftID: rate.GiftID,
	}
	data, err := t.RateRepository.Save(ratemModel)

	helper.ErrorPanic(err)

	rateResponse := response.RateResponse{
		Id: int(data.ID),
	}

	return rateResponse
}

// Delete implements RateService.
func (t *RateServiceImpl) Delete(rateId int) {
	t.RateRepository.Delete(rateId)
}

// FindAll implements RateService.
func (t *RateServiceImpl) FindAll() []response.RateResponse {
	result := t.RateRepository.FindAll()

	var rates []response.RateResponse
	for _, value := range result {
		gift := response.RateResponse{
			UserID: value.UserID,
			GiftID: value.GiftID,
			Rated:  value.Rate,
		}
		rates = append(rates, gift)
	}

	return rates

}

// FindById implements RateService.
func (t *RateServiceImpl) FindById(rateId int) response.RateResponse {
	rateData, err := t.RateRepository.FindById(rateId)
	helper.ErrorPanic(err)

	rateResponse := response.RateResponse{
		Id:     int(rateData.ID),
		UserID: rateData.UserID,
		GiftID: rateData.GiftID,
		Rated:  rateData.Rate,
	}
	return rateResponse
}

// Update implements RateService.
func (t *RateServiceImpl) Update(rate request.UpdateRateRequest) {
	rateData, err := t.RateRepository.FindById(rate.Id)
	helper.ErrorPanic(err)
	rateData.UserID = rate.UserID
	rateData.GiftID = rate.GiftID
	rateData.Rate = rateData.Rate
	t.RateRepository.Update(rateData)
}
