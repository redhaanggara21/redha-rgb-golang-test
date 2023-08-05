package repository

import (
	"golang-gin/pkg/model"
)

type ReedemRepository interface {
	Save(reedem model.Reedem) (reedems model.Reedem, err error)
	Update(reedem model.Reedem)
	Delete(reedemId int)
	FindById(reedemId int) (model.Reedem, error)
	FindAll() []model.Reedem
}
