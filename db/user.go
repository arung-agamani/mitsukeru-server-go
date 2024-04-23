package db

import (
	"github.com/arung-agamani/mitsukeru-server-go/models"
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

func GetUserByUsername(username string, conn *gorm.DB) (*models.User, error) {
	var user models.User
	get := conn.Where("username = ?", username).First(&user)
	if get.Error != nil {
		return nil, get.Error
	}
	return &user, nil
}
