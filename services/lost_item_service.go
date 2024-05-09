package services

import (
	"fmt"
	"github.com/arung-agamani/mitsukeru-server-go/db"
	"github.com/arung-agamani/mitsukeru-server-go/models"
	"github.com/arung-agamani/mitsukeru-server-go/utils/validator"
	"github.com/google/uuid"
)

type CreateLostItemPayload struct {
	Name            string    `json:"name" validate:"required,max=64"`
	Description     string    `json:"description" validate:"max=512"`
	ItemTypeName    string    `json:"itemTypeName" validate:"required"`
	EventID         uuid.UUID `json:"eventId" validate:"required,uuid4"`
	ReporterName    string    `json:"reporterName" validate:"required,max=256"`
	ReporterContact string    `json:"reporterContact" validate:"required,max=64"`
	Returned        bool      `json:"returned"`
	ImageData       string    `json:"imageData"`
	Type            string    `json:"type"`
}

type UpdateLostItemPayload struct {
	ID              uuid.UUID `json:"id" validate:"required,uuid4"`
	Name            string    `json:"name" validate:"required,max=64"`
	Description     string    `json:"description" validate:"max=512"`
	ItemTypeName    string    `json:"itemTypeID" validate:"required"`
	ReporterName    string    `json:"reporterName" validate:"required,max=256"`
	ReporterContact string    `json:"reporterContact" validate:"required,max=64"`
	Returned        bool      `json:"returned"`
}

type LostItemService interface {
	CreateLostItem(p CreateLostItemPayload) (*models.LostItem, error)
	GetLostItem(id uuid.UUID) (*models.LostItem, error)
	ListLostItem() ([]models.LostItem, error)
	UpdateLostItem(p UpdateLostItemPayload) (*models.LostItem, error)
	RemoveLostItem(id uuid.UUID) (bool, error)
}

type lostItemService struct{}

func NewLostItemService() LostItemService { return &lostItemService{} }

func (lis *lostItemService) CreateLostItem(payload CreateLostItemPayload) (*models.LostItem, error) {
	conn := db.GetDB()
	event, err := db.GetEvent(payload.EventID, conn)
	if err != nil {
		return nil, err
	}
	itemType, err := db.GetItemType(payload.ItemTypeName, conn)
	if err != nil {
		return nil, err
	}
	objectName := fmt.Sprintf("%s.png", uuid.New().String())
	if len(payload.ImageData) > 10 {
		_, err = db.S3Upload(payload.ImageData, objectName)
		if err != nil {
			return nil, err
		}
	} else {
		objectName = ""
	}

	li := models.LostItem{
		Item: models.Item{
			Name: payload.Name, Description: payload.Description,
			Event: *event, ItemType: *itemType,
		},
		Type:            payload.Type,
		ReporterName:    payload.ReporterName,
		ReporterContact: payload.ReporterContact,
		Returned:        false,
		Asset:           objectName,
	}
	newLostItem, err := db.CreateLostItem(&li, conn)
	if err != nil {
		return nil, err
	}
	return newLostItem, nil
}

func (lis *lostItemService) GetLostItem(id uuid.UUID) (*models.LostItem, error) {
	if ok, err := validator.ValidateVariable(id, "required,uuid4"); !ok {
		return nil, ValidatorError{
			FieldErrors: *err,
		}
	}
	conn := db.GetDB()
	lostItem, err := db.GetLostItem(id, conn)
	if err != nil {
		return nil, err
	}
	return lostItem, nil
}

func (lis *lostItemService) ListLostItem() ([]models.LostItem, error) {
	conn := db.GetDB()
	lostItems, err := db.ListLostItem(conn)
	if err != nil {
		return nil, err
	}
	return lostItems, nil
}

func (lis *lostItemService) UpdateLostItem(payload UpdateLostItemPayload) (*models.LostItem, error) {
	if ok, err := validator.ValidateStruct(&payload); !ok {
		return nil, ValidatorError{FieldErrors: *err}
	}
	conn := db.GetDB()
	li := models.LostItem{
		Item: models.Item{
			ID:          payload.ID,
			Name:        payload.Name,
			Description: payload.Description,
			ItemTypeID:  payload.ItemTypeName,
		},
		ReporterName:    payload.ReporterName,
		ReporterContact: payload.ReporterContact,
		Returned:        payload.Returned,
	}
	lostItem, err := db.UpdateLostItem(&li, conn)
	if err != nil {
		return nil, err
	}
	return lostItem, nil
}

func (lis *lostItemService) RemoveLostItem(id uuid.UUID) (bool, error) {
	conn := db.GetDB()
	_, err := db.RemoveLostItem(id, conn)
	if err != nil {
		return false, err
	}
	return true, nil
}
