package handlars

import (
	"errors"
	"fmt"
	"log"

	"github.com/ShankaranarayananBR/bougette-backend/cmd/api/requests"
	"github.com/ShankaranarayananBR/bougette-backend/cmd/api/services"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func (h *Handler) RegisterHandler(c echo.Context) error {
	// bind the request body
	payload := new(requests.RegisterUserRequest)
	err := (&echo.DefaultBinder{}).Bind(&payload, c)
	if err != nil {
		return SendBadRequestResponse(c, err.Error())
	}
	validationErrors := h.ValidateBodyRequest(c, *payload)
	if validationErrors != nil {
		fmt.Println(&validationErrors)
		return SendFailedValidationResponse(c, validationErrors)
	}
	fmt.Println(&validationErrors)
	userService := services.NewUserService(h.DB)
	// Check if email exists
	email, err := userService.GetUserByEmail(payload.Email)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("user is present")
	}
	fmt.Println(email)
	// Create user
	newuser, err := userService.RegisterUser(*payload)
	if err != nil {
		return err
	}
	log.Printf("payload:%v", *payload)
	// update user
	return SendSuccessResponse(c, "User regitration successful", newuser)
}
