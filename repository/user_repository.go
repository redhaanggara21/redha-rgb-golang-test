package repository

import (
	"golang-crud-gin/model"
)

type UserRepository interface {
	Save(user model.User)
	Update(user model.User)
	Delete(userId int)
	FindById(userId int) (user model.User, err error)
	FindByEmail(email string) (user model.User, err error)
	FindAll() []model.User
}
