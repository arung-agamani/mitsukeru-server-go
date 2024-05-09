package db

import (
	"errors"
	"github.com/arung-agamani/mitsukeru-server-go/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetEvent(id uuid.UUID, conn *gorm.DB) (*models.Event, error) {
	event := models.Event{
		ID: id,
	}
	get := conn.First(&event)
	if ok, err := HandleError(get.Error); !ok {
		return nil, errors.New(err)
	}
	return &event, nil
}

func ListEvent(conn *gorm.DB) ([]models.Event, error) {
	var events []models.Event
	res := conn.Find(&events)
	if ok, err := HandleError(res.Error); !ok {
		return nil, errors.New(err)
	}
	return events, nil
}

func UpdateEvent(e *models.Event, conn *gorm.DB) (*models.Event, error) {
	i, err := GetEvent(e.ID, conn)
	if err != nil {
		return nil, err
	}
	i.Name = e.Name
	i.Description = e.Description
	i.StartDate = e.StartDate
	i.EndDate = e.EndDate
	i.Location = e.Location
	save := conn.Save(&i)
	if ok, err := HandleError(save.Error); !ok {
		return nil, errors.New(err)
	}
	return i, nil
}

func RemoveEvent(id uuid.UUID, conn *gorm.DB) (bool, error) {
	e, err := GetEvent(id, conn)
	if err != nil {
		return false, err
	}
	del := conn.Delete(e)
	if ok, err := HandleError(del.Error); !ok {
		return false, errors.New(err)
	}
	return true, nil
}
