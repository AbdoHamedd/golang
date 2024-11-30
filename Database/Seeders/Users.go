package Seeders

import (
	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/gorm"
	"project1/Models"
)

func SeedUser(DB *gorm.DB) {
	DB.Create(&Models.User{
		UserName: gofakeit.Name(),
		Email:    gofakeit.Email(),
		Password: gofakeit.Password(true, true, true, true, true, 10),
		Group:    "admin",
		Token:    gofakeit.Email(),
	})
}
