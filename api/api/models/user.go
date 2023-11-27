package models

import "github.com/go-playground/validator/v10"

type User struct {
	Id       int    `json:"Id"`
	Name     string `json:"Name" validate:"required"`
	Email    string `json:"Email" validate:"required"`
	Password string `json:"Password" validate:"required"`
}

func (u *User) UserValidate() error {
	validate := validator.New()
	return validate.Struct(u)
}
