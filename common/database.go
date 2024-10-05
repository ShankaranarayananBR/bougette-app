package common

import (
	"fmt"
	"log"
	"os"

	"github.com/ShankaranarayananBR/bougette-backend/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysql() (*gorm.DB, error) {
	var usermodel models.UserModel
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_NAME")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)
	//fmt.Printf("Printing the dsn:%v", dsn)
	// Open mysql connection
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(usermodel)
	log.Default().Printf("Database connection successful%v", &db)
	return db, nil
}
