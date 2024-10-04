package main

import "github.com/ShankaranarayananBR/bougette-backend/cmd/api/handlars"

func (app *Application) routes(handler handlars.Handler) {
	app.server.GET("/", handler.Healthcheck)
	app.server.POST("/register", handler.RegisterHandler)
}
