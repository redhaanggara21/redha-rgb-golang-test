package repository

import (
	"errors"
	"golang-crud-gin/data/request"
	"golang-crud-gin/helper"
	"golang-crud-gin/model"

	"gorm.io/gorm"
)

type RateRepositoryImpl struct {
	Db *gorm.DB
}

func NewRateREpositoryImpl(Db *gorm.DB) RateRepository {
	return &RateRepositoryImpl{Db: Db}
}

// Delete implements RateRepository.
func (t *RateRepositoryImpl) Delete(rateId int) {
	var rates model.Rating
	result := t.Db.Where("id = ?", rateId).Delete(&rates)
	helper.ErrorPanic(result.Error)
}

// FindAll implements RateRepository.
func (t *RateRepositoryImpl) FindAll() []model.Rating {
	var rates []model.Rating
	result := t.Db.Find(&rates)
	helper.ErrorPanic(result.Error)
	return rates
}

// FindById implements RateRepository.
func (t *RateRepositoryImpl) FindById(rateId int) (model.Rating, error) {
	var rate model.Rating
	result := t.Db.Find(&rate, rateId)
	if result != nil {
		return rate, nil
	} else {
		return rate, errors.New("rate is not found")
	}
}

// Save implements RateRepository.
func (t *RateRepositoryImpl) Save(rate model.Rating) (rates model.Rating, err error) {
	result := t.Db.Create(&rate)
	helper.ErrorPanic(result.Error)
	return rate, nil
}

// Update implements RateRepository.
func (t *RateRepositoryImpl) Update(rate model.Rating) {
	var updateRate = request.UpdateRateRequest{
		Id: int(rate.ID),
	}
	result := t.Db.Model(&rate).Updates(updateRate)
	helper.ErrorPanic(result.Error)
}
