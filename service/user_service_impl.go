package service

import (
	"golang-crud-gin/data/request"
	"golang-crud-gin/data/response"
	"golang-crud-gin/helper"
	"golang-crud-gin/model"
	"golang-crud-gin/repository"

	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	Validate       *validator.Validate
}

func NewUserServiceImpl(userRepository repository.UserRepository, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		Validate:       validate,
	}
}

// Create implements UserService
func (t *UserServiceImpl) Create(user request.CreateUserRequest) {
	err := t.Validate.Struct(user)
	helper.ErrorPanic(err)
	userModel := model.User{
		Email:    user.Email,
		Name:     user.Name,
		Password: user.Password,
		Role:     user.Role,
	}
	t.UserRepository.Save(userModel)
}

// Delete implements TagsService
func (t *UserServiceImpl) Delete(userId int) {
	t.UserRepository.Delete(userId)
}

// FindAll implements TagsService
func (t *UserServiceImpl) FindAll() []response.UserResponse {
	result := t.UserRepository.FindAll()

	var users []response.UserResponse
	for _, value := range result {
		user := response.UserResponse{
			Id:    int(value.ID),
			Name:  value.Name,
			Email: value.Email,
			Role:  value.Role,
		}
		users = append(users, user)
	}

	return users
}

// FindById implements TagsService
func (t *UserServiceImpl) FindById(userId int) response.UserResponse {
	userData, err := t.UserRepository.FindById(userId)
	helper.ErrorPanic(err)

	userResponse := response.UserResponse{
		Id:    int(userData.ID),
		Name:  userData.Name,
		Email: userData.Email,
		Role:  userData.Role,
	}
	return userResponse
}

// FindByEmail implements UserService.
func (t *UserServiceImpl) FindByEmail(email string) response.UserResponse {
	userData, err := t.UserRepository.FindByEmail(email)
	helper.ErrorPanic(err)

	userResponse := response.UserResponse{
		Id:    int(userData.ID),
		Name:  userData.Name,
		Email: userData.Email,
		Role:  userData.Role,
	}
	return userResponse
}

// Update implements TagsService
func (t *UserServiceImpl) Update(users request.UpdateUserRequest) {
	userData, err := t.UserRepository.FindById(users.Id)
	helper.ErrorPanic(err)
	userData.Name = users.Name
	userData.Email = users.Name
	userData.Password = users.Password
	t.UserRepository.Update(userData)
}
