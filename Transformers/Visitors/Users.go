package Visitors

import (
	"project1/Models"
)

func UserTransformer(User Models.User) map[string]interface{} {
	m := make(map[string]interface{})
	m["User_Name"] = User.UserName
	m["email"] = User.Email
	m["token"] = User.Token
	return m
}

func UsersTransformer(Users []Models.User) []map[string]interface{} {
	m := make([]map[string]interface{}, 0)
	for _, user := range Users {
		m = append(m, UserTransformer(user))
	}
	return m
}
