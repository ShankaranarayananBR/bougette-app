package handlars

import (
	"github.com/labstack/echo"
)

// Healthcheck function checks the health of the application
func (h *Handler) Healthcheck(c echo.Context) error {
	healthcheckstruct := struct {
		Health bool `json:"health"`
	}{
		Health: true,
	}
	return c.JSON(200, healthcheckstruct)
}
