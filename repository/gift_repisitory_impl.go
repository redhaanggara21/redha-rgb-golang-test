package repository

import (
	"errors"
	"golang-crud-gin/data/request"
	"golang-crud-gin/helper"
	"golang-crud-gin/model"

	"gorm.io/gorm"
)

type GiftRepositoryImpl struct {
	Db *gorm.DB
}

func NewGiftREpositoryImpl(Db *gorm.DB) GiftRepository {
	return &GiftRepositoryImpl{Db: Db}
}

// Delete implements TagsRepository
func (t *GiftRepositoryImpl) Delete(giftId int) {
	var gifts model.Gift
	result := t.Db.Where("id = ?", giftId).Delete(&gifts)
	helper.ErrorPanic(result.Error)
}

// FindAll implements TagsRepository
func (t *GiftRepositoryImpl) FindAll() []model.Gift {
	var gifts []model.Gift
	result := t.Db.Find(&gifts)
	helper.ErrorPanic(result.Error)
	return gifts
}

// FindById implements TagsRepository
func (t *GiftRepositoryImpl) FindById(giftId int) (tags model.Gift, err error) {
	var gift model.Gift
	result := t.Db.Find(&gift, giftId)
	if result != nil {
		return gift, nil
	} else {
		return gift, errors.New("gift is not found")
	}
}

// Save implements TagsRepository
func (t *GiftRepositoryImpl) Save(gift model.Gift) (model.Gift, error) {
	result := t.Db.Create(&gift)
	helper.ErrorPanic(result.Error)

	return gift, nil
}

// Update implements TagsRepository
func (t *GiftRepositoryImpl) Update(gift model.Gift) {
	var updateGift = request.UpdateGiftRequest{
		Id:          int(gift.ID),
		Name:        gift.Name,
		Description: gift.Description,
		Stock:       gift.Stock,
	}
	result := t.Db.Model(&gift).Updates(updateGift)
	helper.ErrorPanic(result.Error)
}
