package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Drug struct {
	Id           int       `json:"Id"`
	Name         string    `json:"Name" validate:"required"`
	Approved     bool      `json:"Approved" validate:"required"`
	Min_dose     int       `json:"Min_dose" validate:"required"`
	Max_dose     int       `json:"Max_dose" validate:"required"`
	Available_at time.Time `json:"Available_at" validate:"required"`
}

func (u *Drug) DrugValidate() error {
	validate := validator.New()
	return validate.Struct(u)
}
