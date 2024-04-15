package db

import (
	"github.com/arung-agamani/mitsukeru-go/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetUser(id uuid.UUID, conn *gorm.DB) (*models.User, error) {
	user := models.User{ID: id}
	get := conn.First(&user)
	if get.Error != nil {
		return nil, get.Error
	}
	return &user, nil
}
