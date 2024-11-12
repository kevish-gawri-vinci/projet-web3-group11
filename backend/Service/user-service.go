package service

import (
	entity "backend/Entity"
	request "backend/Request"
)

type UserService interface {
	AddUser(request.AddUserRequest) entity.User
	FindById(id int) entity.User
}

type userService struct {
}

// AddUser implements UserService.
func (u *userService) AddUser(addedUser request.AddUserRequest) entity.User {
	print("Called adduser")
	newId := 1
	var user entity.User = entity.User{Id: newId, Username: addedUser.Username, Password: addedUser.Password, IsAdmin: false}
	return user
}

// FindById implements UserService.
func (u *userService) FindById(id int) entity.User {
	panic("unimplemented")
}

func NewUserService() UserService {
	return &userService{}
}
