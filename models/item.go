package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ItemTypeID  string    `json:"typeID"`
}

type ItemType struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
}

type LostItem struct {
	Item
	ReporterName    string `json:"reporterName"`
	ReporterContact string `json:"reporterContact"`
	Returned        bool   `json:"returned"`
}

type FoundItem struct {
	Item
	ReporterName    string `json:"reporterName"`
	ReporterContact string `json:"reporterContact"`
	Returned        bool   `json:"returned"`
}

type DepositItem struct {
	Item
	OwnerName    string `json:"ownerName"`
	OwnerContact string `json:"ownerContact"`
	SlotNumber   int    `json:"slotNumber"`
}

func NewItem() Item           { return Item{} }
func NewLostItem() LostItem   { return LostItem{} }
func NewFoundItem() FoundItem { return FoundItem{} }
