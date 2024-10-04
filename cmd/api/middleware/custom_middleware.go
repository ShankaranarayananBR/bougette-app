package app_middleware

import (
	"fmt"

	"github.com/labstack/echo"
)

func CustomMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	fmt.Println("we are in the custom middleware....")
	return func(ctx echo.Context) error {
		ctx.Response().Header().Set(echo.HeaderServer, "Echo/3.0")
		return next(ctx)
	}
}
