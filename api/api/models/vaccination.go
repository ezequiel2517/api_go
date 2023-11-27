package models

import (
	"time"
)

type Vaccination struct {
	Id      int       `json:"Id"`
	Name    string    `json:"Name" validate:"required"`
	Drug_id int       `json:"Drug_id" validate:"required"`
	Dose    int       `json:"Dose" validate:"required"`
	Date    time.Time `json:"Date" validate:"required"`
}
