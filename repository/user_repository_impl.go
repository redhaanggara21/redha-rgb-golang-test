package repository

import (
	"errors"
	"golang-crud-gin/data/request"
	"golang-crud-gin/helper"
	"golang-crud-gin/model"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func NewUserREpositoryImpl(Db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{Db: Db}
}

// Delete implements UserRepository
func (t *UserRepositoryImpl) Delete(userId int) {
	var users model.User
	result := t.Db.Where("id = ?", userId).Delete(&users)
	helper.ErrorPanic(result.Error)
}

// FindAll implements TagsRepository
func (t *UserRepositoryImpl) FindAll() []model.User {
	var users []model.User
	result := t.Db.Find(&users)
	helper.ErrorPanic(result.Error)
	return users
}

// // FindById implements UserRepository
func (t *UserRepositoryImpl) FindById(userId int) (users model.User, err error) {
	var user model.User
	result := t.Db.Find(&user, userId)
	if result != nil {
		return user, nil
	} else {
		return user, errors.New("user is not found")
	}
}

// Find By Email implements UserRepository.
func (t *UserRepositoryImpl) FindByEmail(email string) (users model.User, err error) {
	var user model.User
	result := t.Db.Select([]string{"id", "name", "email", "role", "password"}).First(&user, "email = ?", email)
	if result != nil {
		return user, nil
	} else {
		return user, errors.New("user is not found")
	}
}

// // Save implements UsersRepository
func (t *UserRepositoryImpl) Save(users model.User) {
	result := t.Db.Create(&users)
	helper.ErrorPanic(result.Error)
}

// Update implements UserRepositoryImpl
func (t *UserRepositoryImpl) Update(users model.User) {
	var updateUser = request.UpdateUserRequest{
		Id:       int(users.ID),
		Name:     users.Name,
		Email:    users.Email,
		Password: users.Password,
		Role:     users.Role,
	}
	result := t.Db.Model(&users).Updates(updateUser)
	helper.ErrorPanic(result.Error)
}
