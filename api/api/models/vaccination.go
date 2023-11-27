package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Vaccination struct {
	Id      int       `json:"Id"`
	Name    string    `json:"Name" validate:"required"`
	Drug_id int       `json:"Drug_id" validate:"required"`
	Dose    int       `json:"Dose" validate:"required"`
	Date    time.Time `json:"Date" validate:"required"`
}

func (u *Vaccination) VaccinationValidate() error {
	validate := validator.New()
	return validate.Struct(u)
}
