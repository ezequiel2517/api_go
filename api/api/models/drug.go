package models

import (
	"time"
)

type Drug struct {
	ID          int       `json:"ID"`
	Name        string    `json:"Name" validate:"required"`
	Approved    bool      `json:"Approved" validate:"required"`
	MinDose     int       `json:"MinDose" validate:"required"`
	MaxDose     int       `json:"MaxDose" validate:"required"`
	AvailableAt time.Time `json:"AvailableAt" validate:"required"`
}
