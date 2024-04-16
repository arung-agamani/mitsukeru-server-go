package services

import (
	"errors"
	"github.com/arung-agamani/mitsukeru-go/db"
	"github.com/arung-agamani/mitsukeru-go/models"
	"github.com/arung-agamani/mitsukeru-go/utils/validator"
)

type CreateItemTypePayload struct {
	Name        string `validate:"required"`
	Description string
}

type UpdateItemTypePayload struct {
	Name        string `validate:"required"`
	Description string
}

type ItemTypeService interface {
	CreateItemType(name, description string) (*models.ItemType, error)
	GetItemType(name string) (*models.ItemType, error)
	ListItemType() ([]models.ItemType, error)
	UpdateItemType(name, description string) (*models.ItemType, error)
	RemoveItemType(name string) (bool, error)
}

type itemTypeService struct{}

func NewItemTypeService() ItemTypeService { return &itemTypeService{} }

func (its *itemTypeService) CreateItemType(name, description string) (*models.ItemType, error) {
	if ok, err := validator.ValidateStruct(&CreateItemTypePayload{
		Name: name, Description: description,
	}); !ok {
		return nil, ValidatorError{
			FieldErrors: *err,
		}
	}
	itemType := models.ItemType{
		Name:        name,
		Description: description,
	}
	dbConn := db.GetDB()
	res := dbConn.Create(&itemType)
	if ok, msg := db.HandleError(res.Error); !ok {
		return nil, errors.New(msg)
	}
	return &itemType, nil
}

func (its *itemTypeService) GetItemType(name string) (*models.ItemType, error) {
	if ok, err := validator.ValidateVariable(&name, "required,alphanum"); !ok {
		return nil, ValidatorError{FieldErrors: *err}
	}
	dbConn := db.GetDB()
	itemType, err := db.GetItemType(name, dbConn)
	if err != nil {
		return nil, err
	}
	return itemType, nil
}

func (its *itemTypeService) ListItemType() ([]models.ItemType, error) {
	dbConn := db.GetDB()
	itemTypes, err := db.ListItemType(dbConn)
	if err != nil {
		return nil, err
	}
	return itemTypes, nil
}

func (its *itemTypeService) UpdateItemType(name, description string) (*models.ItemType, error) {
	if ok, err := validator.ValidateStruct(&UpdateItemTypePayload{
		Name:        name,
		Description: description,
	}); !ok {
		return nil, ValidatorError{FieldErrors: *err}
	}
	dbConn := db.GetDB()
	itemType, err := db.UpdateItemType(name, description, dbConn)
	if err != nil {
		return nil, err
	}
	return itemType, nil

}
func (its *itemTypeService) RemoveItemType(name string) (bool, error) {
	if ok, err := validator.ValidateVariable(name, "required,alphanum"); !ok {
		return false, ValidatorError{
			FieldErrors: *err,
		}
	}
	dbConn := db.GetDB()
	_, err := db.DeleteItemType(name, dbConn)
	if err != nil {
		return false, err
	}
	return true, nil

}
