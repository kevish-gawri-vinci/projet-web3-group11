package service

import (
	entity "backend/Entity"
	request "backend/Request"

	"gorm.io/gorm"
)

type UserService interface {
	AddUser(request.AddUserRequest) entity.User
	FindById(id int) entity.User
}

type userService struct {
	DB *gorm.DB
}

// AddUser implements UserService.
func (u *userService) AddUser(addedUser request.AddUserRequest) entity.User {
	print("Called adduser")
	db := u.DB
	var user entity.User
	db.First(&user)
	return user
}

// FindById implements UserService.
func (u *userService) FindById(id int) entity.User {
	panic("unimplemented")
}

func NewUserService(db *gorm.DB) UserService {
	return &userService{DB: db}
}
