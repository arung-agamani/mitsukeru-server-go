package db

import (
	"errors"
	"github.com/arung-agamani/mitsukeru-server-go/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateItem(item *models.Item, conn *gorm.DB) (*models.Item, error) {
	create := conn.Create(item)
	if ok, err := HandleError(create.Error); !ok {
		return nil, errors.New(err)
	}
	return item, nil
}
func GetItem(id uuid.UUID, conn *gorm.DB) (*models.Item, error) {
	item := models.Item{ID: id}
	get := conn.First(&item)
	if ok, err := HandleError(get.Error); !ok {
		return nil, errors.New(err)
	}
	return &item, nil
}
func ListItems(conn *gorm.DB) ([]models.Item, error) {
	var items []models.Item
	res := conn.Find(&items)
	if ok, err := HandleError(res.Error); !ok {
		return nil, errors.New(err)
	}
	return items, nil
}
func UpdateItem(item *models.Item, conn *gorm.DB) (*models.Item, error) {
	i, err := GetItem(item.ID, conn)
	if err != nil {
		return nil, err
	}
	i.Name = item.Name
	i.Description = item.Description
	i.ItemType = item.ItemType
	i.Event = item.Event
	res := conn.Save(&i)
	if ok, err := HandleError(res.Error); !ok {
		return nil, errors.New(err)
	}
	return i, nil
}

func DeleteItem(id uuid.UUID, conn *gorm.DB) (bool, error) {
	i, err := GetItem(id, conn)
	if err != nil {
		return false, err
	}
	del := conn.Delete(i)
	if ok, err := HandleError(del.Error); !ok {
		return false, errors.New(err)
	}
	return true, nil
}
