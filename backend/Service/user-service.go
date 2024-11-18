package service

import (
	entity "backend/Entity"
	request "backend/Request"
	utils "backend/Utils"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	AddUser(request.UserRequest) (entity.User, error)
	Login(request.UserRequest) (entity.User, error, string)
}

type userService struct {
	DB *gorm.DB
}

// AddUser implements UserService.
func (u *userService) AddUser(addedUser request.UserRequest) (entity.User, error) {
	print("Called adduser")
	db := u.DB
	var user entity.User

	if err := db.Create(&entity.User{
		Username: addedUser.Username,
		Password: addedUser.Password,
	}).Error; err != nil {
		return entity.User{}, err
	}

	return user, nil
}

// Login implements UserService.
func (u *userService) Login(requestInput request.UserRequest) (entity.User, error, string) {
	db := u.DB
	var user entity.User
	result := db.First(&user, "username = ?", requestInput.Username)

	if result.RowsAffected != 1 {
		return user, errors.New("No user found with this username"), ""
	}

	//Check if the password is correct or not
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(requestInput.Password)); err != nil {
		return entity.User{}, errors.New("incorrect password"), ""
	}

	//Generate JWT Token
	token, err := utils.CreateToken(user.Username)

	if err != nil {
		return user, errors.New("Error during the generation of JWT"), ""
	}

	return user, nil, token
}

func NewUserService(db *gorm.DB) UserService {
	return &userService{DB: db}
}
