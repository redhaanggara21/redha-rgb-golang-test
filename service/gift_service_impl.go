package service

import (
	"golang-crud-gin/data/request"
	"golang-crud-gin/data/response"
	"golang-crud-gin/helper"
	"golang-crud-gin/model"
	"golang-crud-gin/repository"

	"github.com/go-playground/validator/v10"
)

type GiftServiceImpl struct {
	GiftRepository repository.GiftRepository
	Validate       *validator.Validate
}

func NewGiftServiceImpl(giftRepository repository.GiftRepository, validate *validator.Validate) GiftService {
	return &GiftServiceImpl{
		GiftRepository: giftRepository,
		Validate:       validate,
	}
}

// Update implements GiftService.
func (t *GiftServiceImpl) Update(gift request.UpdateGiftRequest) {
	giftData, err := t.GiftRepository.FindById(gift.Id)
	helper.ErrorPanic(err)
	giftData.Name = gift.Name
	giftData.Description = gift.Description
	giftData.Stock = gift.Stock
	t.GiftRepository.Update(giftData)
}

// Create implements TagsService
func (t *GiftServiceImpl) Create(gift request.CreateGiftRequest) response.GiftResponse {
	err := t.Validate.Struct(gift)
	helper.ErrorPanic(err)
	giftModel := model.Gift{
		Name:        gift.Name,
		Description: gift.Description,
		Image:       gift.Image,
		Stock:       gift.Stock,
	}
	data, err := t.GiftRepository.Save(giftModel)

	helper.ErrorPanic(err)

	tagResponse := response.GiftResponse{
		Id: int(data.ID),
	}

	return tagResponse
}

// Delete implements TagsService
func (t *GiftServiceImpl) Delete(giftId int) {
	t.GiftRepository.Delete(giftId)
}

// FindAll implements GiftService.
func (t *GiftServiceImpl) FindAll() []response.GiftResponse {
	result := t.GiftRepository.FindAll()

	var gifts []response.GiftResponse
	for _, value := range result {
		gift := response.GiftResponse{
			Id:          int(value.ID),
			Name:        value.Name,
			Description: value.Description,
			Stock:       value.Stock,
		}
		gifts = append(gifts, gift)
	}

	return gifts
}

// FindById implements GiftService.
func (t *GiftServiceImpl) FindById(giftId int) request.UpdateGiftRequest {
	giftData, err := t.GiftRepository.FindById(giftId)
	helper.ErrorPanic(err)

	tagResponse := request.UpdateGiftRequest{
		Id:          int(giftData.ID),
		Name:        giftData.Name,
		Description: giftData.Description,
		Stock:       giftData.Stock,
	}
	return tagResponse
}
