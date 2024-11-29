package Models

import "gorm.io/gorm"

type Users []User

type User struct {
	gorm.Model
	UserName string `json:"user_name" gorm:"type:varchar(200)"`
	Email    string `json:"email" gorm:"type:varchar(200)"`
	Password string `json:"password" gorm:"type:varchar(200)"`
	Token    string `json:"token" gorm:"type:varchar(200)"`
	Group    string `json:"Group" gorm:"type:varchar(200)"`
}

//// another aproch for Transformer
//func (user User) Transformer() map[string]interface{} {
//	m := make(map[string]interface{})
//	m["User_Name"] = user.UserName
//	m["email"] = user.Email
//	m["token"] = user.Token
//	return m
//}
//func (users Users) Transformer() []map[string]interface{} {
//	m := make([]map[string]interface{}, 0)
//	for _, user := range users {
//		m = append(m, user.Transformer())
//	}
//	return m
//}
