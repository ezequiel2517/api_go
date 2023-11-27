package models

import (
	"time"
)

type Drug struct {
	Id           int       `json:"Id"`
	Name         string    `json:"Name" validate:"required"`
	Approved     bool      `json:"Approved" validate:"required"`
	Min_dose     int       `json:"Min_dose" validate:"required"`
	Max_dose     int       `json:"Max_dose" validate:"required"`
	Available_at time.Time `json:"Available_at" validate:"required"`
}
