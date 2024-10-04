package main

import (
	"github.com/ShankaranarayananBR/bougette-backend/common"
	"github.com/ShankaranarayananBR/bougette-backend/internal/models"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error while loading .env file")
	}
	db, err := common.NewMysql()
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&models.UserModel{})
	if err != nil {
		panic(err)
	}

}
