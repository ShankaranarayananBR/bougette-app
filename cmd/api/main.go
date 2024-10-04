package main

import (
	"fmt"
    "log"
    "os"

	"github.com/ShankaranarayananBR/bougette-backend/cmd/api/handlars"
	app_middleware "github.com/ShankaranarayananBR/bougette-backend/cmd/api/middleware"
	"github.com/ShankaranarayananBR/bougette-backend/common"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Application struct {
	logger  echo.Logger
	server  *echo.Echo
	handler handlars.Handler
}

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("Error while loading .env file")
	}

}
func main() {
	e := echo.New()
	db, err := common.NewMysql()
	if err != nil {
		panic("Error while connecting to database")
	}
	h := handlars.Handler{
		DB: db,
	}
	app := Application{
		logger:  e.Logger,
		server:  e,
		handler: h,
	}
	log.Printf("")
	e.Use(middleware.Logger(), app_middleware.CustomMiddleware)
	fmt.Println(app)
	app.routes(h)
	e.Logger.Fatal(e.Start(os.Getenv("APP_PORT")))
}
