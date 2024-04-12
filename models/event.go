package models

import (
	"gorm.io/gorm"
	"time"
)

type Event struct {
	gorm.Model
	Name        string    `json:"name"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"startDate"`
	EndDate     time.Time `json:"endDate"`
	Location    string    `json:"location"`
}

func NewEventType() Event {
	return Event{}
}
