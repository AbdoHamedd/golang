package Visitors

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"project1/Models"
	"project1/Validations"
)

func RegisterValidation(User Models.User) validation.Errors {
	return validation.Errors{
		"name":     validation.Validate(User.UserName, Validations.RequiredRule(), Validations.Min_MaxRule()),
		"email":    validation.Validate(User.Email, Validations.RequiredRule(), Validations.IsEmailRule(), Validations.Min_MaxRule()),
		"password": validation.Validate(User.Password, Validations.RequiredRule(), Validations.Min_MaxRule()),
	}
}
