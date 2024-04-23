package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ImageAsset struct {
	gorm.Model
	ID  uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	URL string    `json:"url"`
}
