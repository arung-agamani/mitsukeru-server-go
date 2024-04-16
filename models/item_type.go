package models

import "gorm.io/gorm"

type ItemType struct {
	gorm.Model
	Name        string `json:"name" gorm:"unique"`
	Description string `json:"description"`
}
