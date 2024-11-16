package service

import (
	entity "backend/Entity"
	request "backend/Request"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	AddUser(request.UserRequest) entity.User
	Login(request.UserRequest) entity.User
}

type userService struct {
	DB *gorm.DB
}

// AddUser implements UserService.
func (u *userService) AddUser(addedUser request.UserRequest) entity.User {
	print("Called adduser")
	db := u.DB
	var user entity.User
	db.First(&user)
	return user
}

// Login implements UserService.
func (u *userService) Login(requestInput request.UserRequest) entity.User {
	db := u.DB
	var user entity.User
	db.First(&user, "username = ?", requestInput.Username)
	println(user.Password)

	//Check if the password is correct or not
	res := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(requestInput.Password))
	println(res == nil)

	return user
}

func NewUserService(db *gorm.DB) UserService {
	return &userService{DB: db}
}
