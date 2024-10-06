package services

import (
	"errors"
	"log"

	"github.com/ShankaranarayananBR/bougette-backend/cmd/api/requests"
	"github.com/ShankaranarayananBR/bougette-backend/internal/models"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

func (usr *UserService) RegisterUser(userrequest requests.RegisterUserRequest) (*models.UserModel, error) {
	createdUser := models.UserModel{
		FirstName: userrequest.FirstName,
		LastName:  userrequest.LastName,
		Email:     userrequest.Email,
		Password:  userrequest.Password,
	}
	log.Printf("createduser:%v", createdUser)
	result := usr.db.Create(&createdUser)
	if result.Error != nil {
		return nil, errors.New("registration failed")
	}

	return &createdUser, nil
}

func (usr *UserService) GetUserByEmail(email string) (*models.UserModel, error) {
	var user models.UserModel
	result := usr.db.Where("email= ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (usr *UserService) UpdateUserDetails(userreqest requests.RegisterUserRequest, detail string) (*models.UserModel, error) {

	return nil, nil
}
