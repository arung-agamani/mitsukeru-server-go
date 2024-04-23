package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID          uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Username    string    `json:"username" validate:"required,alphanum,min=8,max=32" gorm:"unique"`
	DisplayName string    `json:"displayName" validate:"max=64"`
	Bio         string    `json:"bio" validate:"max=256"`
	Password    string    `json:"password" validate:"required,min=8,max=32,alphanumunicode"`
}
