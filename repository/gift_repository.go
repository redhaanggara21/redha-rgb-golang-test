package repository

import (
	"golang-crud-gin/model"
)

type GiftRepository interface {
	Save(gift model.Gift) (gifts model.Gift, err error)
	Update(gift model.Gift)
	Delete(giftId int)
	FindById(giftId int) (model.Gift, error)
	FindAll() []model.Gift
}
