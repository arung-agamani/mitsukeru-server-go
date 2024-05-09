package db

import (
	"errors"
	"github.com/arung-agamani/mitsukeru-server-go/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateLostItem(item *models.LostItem, conn *gorm.DB) (*models.LostItem, error) {
	create := conn.Create(item)
	if ok, err := HandleError(create.Error); !ok {
		return nil, errors.New(err)
	}
	return item, nil
}

func GetLostItem(id uuid.UUID, conn *gorm.DB) (*models.LostItem, error) {
	i := models.LostItem{
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

func ListLostItem(conn *gorm.DB) ([]models.LostItem, error) {
	var lostItems []models.LostItem
	res := conn.Find(&lostItems)
	if ok, err := HandleError(res.Error); !ok {
		return nil, errors.New(err)
	}
	return lostItems, nil
}

func UpdateLostItem(li *models.LostItem, conn *gorm.DB) (*models.LostItem, error) {
	i, err := GetLostItem(li.ID, conn)
	if err != nil {
		return nil, err
	}
	i.Name = li.Name
	i.Description = li.Description
	i.ItemTypeID = li.ItemTypeID
	i.ReporterName = li.ReporterName
	i.ReporterContact = li.ReporterContact
	i.Returned = li.Returned

	save := conn.Save(&i)
	if ok, err := HandleError(save.Error); !ok {
		return nil, errors.New(err)
	}
	return i, nil
}

func RemoveLostItem(id uuid.UUID, conn *gorm.DB) (bool, error) {
	i, err := GetLostItem(id, conn)
	if err != nil {
		return false, err
	}
	del := conn.Delete(i)
	if ok, err := HandleError(del.Error); !ok {
		return false, errors.New(err)
	}
	return true, nil
}
