package models

import "gorm.io/gorm"

// We will have all the models for each DB table

type UserModel struct {
	gorm.Model
	FirstName *string `gorm:"type:varchar(200)"`
	LastName  *string `gorm:"type:varchar(200)"`
	Gender    *string `gorm:"type:varchar(200); not null; unique"`
	Password  string  `gorm:"type:varchar(50)"`
	Email     string  `gorm:"type:varchar(200); not null"`
}

func (u *UserModel) TableName() string {
	return "users"
}
