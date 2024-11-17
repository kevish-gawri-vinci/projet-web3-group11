package service

import (
	entity "backend/Entity"
	request "backend/Request"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	AddUser(request.UserRequest) (entity.User, error)
	Login(request.UserRequest) (entity.User, error)
}

type userService struct {
	DB *gorm.DB
}

// AddUser implements UserService.
func (u *userService) AddUser(addedUser request.UserRequest) (entity.User, error) {
	print("Called adduser")
	db := u.DB
	var user entity.User
	
	if err := db.Create(&Entity.User{
		Username: addedUser.Username,
		Password: addedUser.Password,
	}).Error; err != nil {
		return Entity.User{}, err
	}

	return user, nil
}

// Login implements UserService.
func (u *userService) Login(requestInput request.UserRequest) (entity.User, err) {
	db := u.DB
	var user entity.User
	
	if err := db.First(&user, "username = ?", requestInput.Username).Error; err != nil {
		return Entity.User{}, errors.New("user not found")
	}

	//Check if the password is correct or not
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(requestInput.Password)); err != nil {
		return Entity.User{}, errors.New("incorrect password")
	}

	return user, nil
}

func NewUserService(db *gorm.DB) UserService {
	return &userService{DB: db}
}
