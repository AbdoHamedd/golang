package Database

import (
	"gorm.io/gorm"
	"project1/Models"
)

func Migrate(DB *gorm.DB) {
	DB.AutoMigrate(
		&Models.User{},
	)
}
