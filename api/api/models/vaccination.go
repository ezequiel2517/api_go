package models

import (
	"time"
)

type Vaccination struct {
	ID     int       `json:"ID"`
	Name   string    `json:"Name" validate:"required"`
	DrugId int       `json:"DrugId" validate:"required"`
	Dose   int       `json:"Dose" validate:"required"`
	Fecha  time.Time `json:"Fecha" validate:"required"`
}
