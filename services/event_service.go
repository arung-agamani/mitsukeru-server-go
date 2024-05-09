package services

import (
	"errors"
	"github.com/arung-agamani/mitsukeru-server-go/db"
	"github.com/arung-agamani/mitsukeru-server-go/models"
	"github.com/arung-agamani/mitsukeru-server-go/utils/validator"
	"github.com/google/uuid"
	"time"
)

type CreateEventPayload struct {
	Name        string    `json:"name" validate:"required"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"startDate" validate:"required"`
	EndDate     time.Time `json:"endDate" validate:"required"`
	Location    string    `json:"location" validate:"required"`
}

type UpdateEventPayload struct {
	ID          uuid.UUID `json:"id" validate:"required"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"startDate"`
	EndDate     time.Time `json:"endDate"`
	Location    string    `json:"location"`
}

type EventService interface {
	CreateEvent(p CreateEventPayload) (*models.Event, error)
	GetEvent(id uuid.UUID) (*models.Event, error)
	ListEvent() ([]models.Event, error)
	UpdateEvent(p UpdateEventPayload) (*models.Event, error)
	RemoveEvent(id uuid.UUID) (bool, error)
}

type eventService struct{}

func NewEventService() EventService { return &eventService{} }

func (ev *eventService) CreateEvent(p CreateEventPayload) (*models.Event, error) {
	if ok, err := validator.ValidateStruct(&p); !ok {
		return nil, ValidatorError{FieldErrors: *err}
	}
	dbConn := db.GetDB()
	event := models.Event{
		Name:        p.Name,
		Description: p.Description,
		StartDate:   p.StartDate,
		EndDate:     p.EndDate,
		Location:    p.Location,
	}
	createRes := dbConn.Create(&event)
	if ok, msg := db.HandleError(createRes.Error); !ok {
		return nil, errors.New(msg)
	}
	return &event, nil
}

func (ev *eventService) GetEvent(id uuid.UUID) (*models.Event, error) {
	var event models.Event = models.Event{ID: id}
	dbConn := db.GetDB()
	getRes := dbConn.Select("name", "description", "start_date", "end_date", "location").First(&event)
	if ok, msg := db.HandleError(getRes.Error); !ok {
		return nil, errors.New(msg)
	}
	return &event, nil
}

func (ev *eventService) ListEvent() ([]models.Event, error) {
	conn := db.GetDB()
	events, err := db.ListEvent(conn)
	if err != nil {
		return nil, err
	}
	return events, nil
}
func (ev *eventService) UpdateEvent(p UpdateEventPayload) (*models.Event, error) {
	if ok, err := validator.ValidateStruct(&p); !ok {
		return nil, ValidatorError{FieldErrors: *err}
	}
	conn := db.GetDB()
	e := models.Event{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Location:    p.Location,
		StartDate:   p.StartDate,
		EndDate:     p.EndDate,
	}
	event, err := db.UpdateEvent(&e, conn)
	if err != nil {
		return nil, err
	}
	return event, nil
}

func (ev *eventService) RemoveEvent(id uuid.UUID) (bool, error) {
	conn := db.GetDB()
	_, err := db.RemoveEvent(id, conn)
	if err != nil {
		return false, err
	}
	return true, nil
}
