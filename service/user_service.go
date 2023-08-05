package service

import (
	"golang-crud-gin/data/request"
	"golang-crud-gin/data/response"
)

type UserService interface {
	Create(user request.CreateUserRequest)
	Update(user request.UpdateUserRequest)
	Delete(userId int)
	FindById(userId int) response.UserResponse
	FindByEmail(email string) response.UserResponse
	FindAll() []response.UserResponse
}
