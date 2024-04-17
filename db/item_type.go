package db

import (
	"errors"
	"github.com/arung-agamani/mitsukeru-server-go/models"
	"gorm.io/gorm"
)

// GetItemType Get ItemType with given (unique) name
// Parameters :
//
//	name string: Name of ItemType to retrieve
//	conn *gorm.DB: An instance to connected gorm.DB instance
//
// Returns:
//
//	*model.ItemType: Returns a reference to retrieved ItemType
//	error: Returns error object if operation fails
func GetItemType(name string, conn *gorm.DB) (*models.ItemType, error) {
	var itemType models.ItemType
	get := conn.Where("name = ?", name).First(&itemType)
	if get.Error != nil {
		return nil, get.Error
	}
	return &itemType, nil
}

// UpdateItemType Updates ItemType description
// Parameters :
//
//	name string: Name of ItemType to modify
//	description string: Description of ItemType that will be modified
//	conn *gorm.DB: An instance to connected gorm.DB instance
//
// Returns:
//
//	*model.ItemType: Returns a reference to updated ItemType
//	error: Returns error object if operation fails
func UpdateItemType(name, description string, conn *gorm.DB) (*models.ItemType, error) {
	itemType, err := GetItemType(name, conn)
	if err != nil {
		return nil, err
	}
	itemType.Description = description
	res := conn.Save(&itemType)
	if ok, err := HandleError(res.Error); !ok {
		return itemType, errors.New(err)
	}
	return itemType, nil
}

// ListItemType Get a list of all item types in database
// Parameters :
//
//	conn *gorm.DB: An instance to connected gorm.DB instance
//
// Returns:
//
//	[]models.ItemType: A list of ItemType object
//	error: Returns object if operation fails
func ListItemType(conn *gorm.DB) ([]models.ItemType, error) {
	var itemTypes []models.ItemType
	res := conn.Find(&itemTypes)
	if ok, err := HandleError(res.Error); !ok {
		return nil, errors.New(err)
	}
	return itemTypes, nil
}

// DeleteItemType Delete item type
// Parameters:
//
//	name string: The name of item type to delete
//	conn *gorm.DB: An instance to connected gorm.DB instance
//
// Returns:
//
//	bool: Indicates if delete operation succeed or not
//	error: Returns error object if operation fails
func DeleteItemType(name string, conn *gorm.DB) (bool, error) {
	itemType, err := GetItemType(name, conn)
	if err != nil {
		return false, err
	}
	res := conn.Delete(&itemType)
	if ok, err := HandleError(res.Error); !ok {
		return false, errors.New(err)
	}
	return true, nil
}
