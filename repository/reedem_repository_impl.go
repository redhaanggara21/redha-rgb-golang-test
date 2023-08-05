package repository

import (
	"errors"
	"golang-crud-gin/data/request"
	"golang-crud-gin/helper"
	"golang-crud-gin/model"

	"gorm.io/gorm"
)

type ReedemRepositoryImpl struct {
	Db *gorm.DB
}

func NewReedemREpositoryImpl(Db *gorm.DB) ReedemRepository {
	return &ReedemRepositoryImpl{Db: Db}
}

// Delete implements ReedemRepository.
func (t *ReedemRepositoryImpl) Delete(reedemId int) {
	var reedems model.Reedem
	result := t.Db.Where("id = ?", reedemId).Delete(&reedems)
	helper.ErrorPanic(result.Error)
}

// FindAll implements ReedemRepository.
func (t *ReedemRepositoryImpl) FindAll() []model.Reedem {
	var reedems []model.Reedem
	result := t.Db.Find(&reedems)
	helper.ErrorPanic(result.Error)
	return reedems
}

// FindById implements ReedemRepository.
func (t *ReedemRepositoryImpl) FindById(reedemId int) (model.Reedem, error) {
	var reedem model.Reedem
	result := t.Db.Find(&reedem, reedemId)
	if result != nil {
		return reedem, nil
	} else {
		return reedem, errors.New("reedem is not found")
	}
}

// Save implements ReedemRepository.
func (t *ReedemRepositoryImpl) Save(reedem model.Reedem) (reedems model.Reedem, err error) {
	result := t.Db.Create(&reedems)
	helper.ErrorPanic(result.Error)

	return reedem, nil
}

// Update implements ReedemRepository.
func (t *ReedemRepositoryImpl) Update(reedem model.Reedem) {
	var updateReedem = request.UpdateReedemRequest{
		Id:     int(reedem.ID),
		UserID: reedem.UserID,
		GiftID: reedem.GiftID,
		Amount: reedem.Amount,
	}
	result := t.Db.Model(&reedem).Updates(updateReedem)
	helper.ErrorPanic(result.Error)
}
