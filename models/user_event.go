package models

import (
	"github.com/google/uuid"
)

type UserEventRelation struct {
	UserID  uuid.UUID `json:"user" gorm:"primaryKey"`
	EventID uuid.UUID `json:"event" gorm:"primaryKey"`
	Role    string    `json:"role"`
}
