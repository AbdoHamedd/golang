package Models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `json:"user_name" gorm:"type:varchar(200)"`
	Email    string `json:"email" gorm:"type:varchar(200)"`
	Password string `json:"password" gorm:"type:varchar(200)"`
	Token    string `json:"token" gorm:"type:varchar(200)"`
	Group    string `json:"token" gorm:"type:varchar(200)"`
}
