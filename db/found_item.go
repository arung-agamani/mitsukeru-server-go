package db

import (
	"errors"
	"github.com/arung-agamani/mitsukeru-server-go/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetFoundItem(id uuid.UUID, conn *gorm.DB) (*models.FoundItem, error) {
	i := models.FoundItem{
		Item: models.Item{
			ID: id,
		},
	}
	get := conn.First(&i)
	if ok, err := HandleError(get.Error); !ok {
		return nil, errors.New(err)
	}
	return &i, nil
}

func UpdateFoundItem(li *models.FoundItem, conn *gorm.DB) (*models.FoundItem, error) {
	i, err := GetFoundItem(li.ID, conn)
	if err != nil {
		return nil, err
	}
	i.Name = li.Name
	i.Description = li.Description
	i.ItemType = li.ItemType
	i.ReporterName = li.ReporterName
	i.ReporterContact = li.ReporterContact
	i.Returned = li.Returned
	save := conn.Save(&i)
	if ok, err := HandleError(save.Error); !ok {
		return nil, errors.New(err)
	}
	return i, nil
}
