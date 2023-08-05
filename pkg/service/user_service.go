package service

import (
	"golang-gin/pkg/data/request"
	"golang-gin/pkg/data/response"
)

type UserService interface {
	Create(user request.CreateUserRequest)
	Update(user request.UpdateUserRequest)
	Delete(userId int)
	FindById(userId int) response.UserResponse
	FindByEmail(email string) response.UserResponse
	FindAll() []response.UserResponse
}
