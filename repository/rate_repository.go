package repository

import (
	"golang-crud-gin/model"
)

type RateRepository interface {
	Save(rate model.Rating) (rates model.Rating, err error)
	Update(rateId model.Rating)
	Delete(rateId int)
	FindById(rateId int) (model.Rating, error)
	FindAll() []model.Rating
}
