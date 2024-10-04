package handlars

import (
	"net/http"

	"github.com/labstack/echo"
)

type ApiResponse map[string]any

type JSONSuccess struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type JSONFailedValidationResult struct {
	Success bool               `json:"success"`
	Message string             `json:"message"`
	Errors  []*ValidationError `json:"errors"`
}

type JSONErrorResponse struct {
	Success bool             `json:"success"`
	Message string           `json:"message"`
	Errors  *ValidationError `json:"errors"`
}

func SendSuccessResponse(c echo.Context, message string, data interface{}) error {
	return c.JSON(http.StatusOK, JSONSuccess{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func SendFailedValidationResponse(c echo.Context, errors []*ValidationError) error {
	return c.JSON(http.StatusUnprocessableEntity, JSONFailedValidationResult{
		Success: false,
		Errors:  errors,
	})
}

func SendErrorResponse(c echo.Context, statuscode int, message string) error {
	return c.JSON(statuscode, JSONErrorResponse{
		Success: false,
		Message: message,
	})
}

func SendBadRequestResponse(c echo.Context, message string) error {
	return SendErrorResponse(c, http.StatusBadRequest, message)
}

func SendNotFoundResponse(c echo.Context, message string) error {
	return SendErrorResponse(c, http.StatusNotFound, message)
}

func SendInternalServerErrorResponse(c echo.Context, message string) error {
	return SendErrorResponse(c, http.StatusInternalServerError, message)
}
