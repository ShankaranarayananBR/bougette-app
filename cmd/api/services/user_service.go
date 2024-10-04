package services

import (
	"log"

	"github.com/ShankaranarayananBR/bougette-backend/cmd/api/requests"
	"github.com/ShankaranarayananBR/bougette-backend/common"
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
	//Hash the password, should not store
	password, err := common.HashPassword(userrequest.Password)
	if err != nil {
		return nil, err
	}
	log.Printf("Hashed password:%v", password)
	if common.CheckPasswordHash(userrequest.Password, password) {
		log.Println("Password hashed successfully")
	}

	return nil, nil
}

func (usr *UserService) GetUserByEmail(email string) (*models.UserModel, error) {
	var user models.UserModel
	result := usr.db.Where("email= ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
