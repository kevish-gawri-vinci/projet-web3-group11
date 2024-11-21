package service

import (
	entity "backend/Entity"
	request "backend/Request"
	utils "backend/Utils"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	AddUser(request.UserRequest) (entity.User, *utils.ErrorStruct)
	Login(request.UserRequest) (entity.User, string, *utils.ErrorStruct)
	GetUserRole(userId int) (request.UserRoleRequest, *utils.ErrorStruct)
}

type userService struct {
	DB *gorm.DB
}

// AddUser implements UserService.
func (u *userService) AddUser(addedUser request.UserRequest) (entity.User, *utils.ErrorStruct) {
	print("Called adduser")
	db := u.DB
	var user entity.User

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(addedUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return entity.User{}, &utils.ErrorStruct{Msg: "Bcrypt error : " + err.Error(), Code: http.StatusInternalServerError}
	}

	if err := db.Create(&entity.User{
		Username: addedUser.Username,
		Password: string(hashPassword),
	}).Error; err != nil {
		return entity.User{}, &utils.ErrorStruct{Msg: "DB error : " + err.Error(), Code: http.StatusInternalServerError}
	}

	return user, nil
}

// Login implements UserService.
func (u *userService) Login(requestInput request.UserRequest) (entity.User, string, *utils.ErrorStruct) {
	db := u.DB
	var user entity.User
	result := db.First(&user, "username = ?", requestInput.Username)

	if result.RowsAffected != 1 {
		return user, "", &utils.ErrorStruct{Msg: "User not found", Code: http.StatusNotFound}
	}

	//Check if the password is correct or not
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(requestInput.Password)); err != nil {
		return entity.User{}, "", &utils.ErrorStruct{Msg: "Username or incorrect password", Code: http.StatusUnauthorized}
	}

	//Generate JWT Token
	token, err := utils.CreateToken(user.Username, user.ID, user.IsAdmin)

	if err != nil {
		return user, "", err
	}

	return user, token, nil
}

func (u *userService) GetUserRole(userId int) (request.UserRoleRequest, *utils.ErrorStruct) {
	db := u.DB
	user := entity.User{ID: userId}
	result := db.First(&user)
	var response request.UserRoleRequest
	if result.Error != nil || result.RowsAffected == 0 {
		return request.UserRoleRequest{}, &utils.ErrorStruct{Msg: "Error : no user found", Code: http.StatusNotFound}
	}
	response.IsAdmin = user.IsAdmin
	response.Username = user.Username
	return response, nil
}

func NewUserService(db *gorm.DB) UserService {
	return &userService{DB: db}
}
