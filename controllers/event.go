package controllers

import (
	"errors"
	"fmt"
	"github.com/arung-agamani/mitsukeru-server-go/db"
	"github.com/arung-agamani/mitsukeru-server-go/models"
	"github.com/arung-agamani/mitsukeru-server-go/responses"
	"github.com/arung-agamani/mitsukeru-server-go/utils/logger"
	"github.com/arung-agamani/mitsukeru-server-go/utils/parser"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type CreateEventRequestPayload struct {
	Name        string    `json:"name" validate:"required"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"startDate" validate:"required"`
	EndDate     time.Time `json:"endDate" validate:"required"`
	Location    string    `json:"location" validate:"required"`
}

type GetEventRequestPayload struct {
	ID uuid.UUID `json:"id" validate:"required"`
}

type DeleteEventRequestPayload struct {
	ID uuid.UUID `json:"id" validate:"required"`
}

type GetEventResponsePayload struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"startDate"`
	EndDate     time.Time `json:"endDate"`
	Location    string    `json:"location"`
}

type UpdateEventRequestPayload struct {
	ID          uuid.UUID `json:"id" validate:"required"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"startDate"`
	EndDate     time.Time `json:"endDate"`
	Location    string    `json:"location"`
}

type UpdateEventResponsePayload struct {
	GetEventResponsePayload
	ID uuid.UUID `json:"id"`
}

func CreateEventHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload CreateEventRequestPayload
		if parser.ParseJSONPayload(w, r, &payload) != true {
			return
		}
		eventCreatePayload := models.Event{
			Name:        payload.Name,
			Description: payload.Description,
			StartDate:   payload.StartDate,
			EndDate:     payload.EndDate,
			Location:    payload.Location,
		}

		dbConn := db.GetDB()
		createRes := dbConn.Create(&eventCreatePayload)
		if createRes.Error != nil {
			responses.ErrResponse(w, &responses.ErrorResponse{
				Status:  500,
				Message: "Error when creating new event",
				Error:   nil,
			})
			logger.Error("Error when creating new event: %s", createRes.Error.Error())
			return
		}
		responses.OkResponse(w, &responses.Response{
			Status:  201,
			Message: fmt.Sprintf("Event created. Affected rows: %d", createRes.RowsAffected),
			Data:    eventCreatePayload,
		})
	}

}

func DeleteEventHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, _ := uuid.Parse(vars["eventId"])
		payload := DeleteEventRequestPayload{
			ID: id,
		}
		if parser.ParseStructPayload(w, r, &payload) != true {
			return
		}
		dbConn := db.GetDB()
		event := models.Event{ID: payload.ID}
		getRes := dbConn.Select("id").First(&event)
		if getRes.Error != nil {
			var errMsg string = "Error when getting event"
			var status int = 500
			if errors.Is(getRes.Error, gorm.ErrRecordNotFound) {
				errMsg = "Record not found"
				status = 404
			}
			responses.ErrResponse(w, &responses.ErrorResponse{
				Status:  status,
				Message: errMsg,
				Error:   nil,
			})
			logger.Errorf("Error when getting event: %s", getRes.Error.Error())
			return
		}
		delRes := dbConn.Delete(&event)
		if delRes.Error != nil {
			responses.ErrResponse(w, &responses.ErrorResponse{
				Status:  500,
				Message: "Error when deleting event",
				Error:   nil,
			})
			logger.Error(fmt.Sprintf("Error when deleting event: %s", delRes.Error.Error()))
			return
		}
		responses.OkResponse(w, &responses.Response{
			Status:  200,
			Message: "Event has been deleted",
			Data:    nil,
		})
	}
}

func GetEventHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, _ := uuid.Parse(vars["eventId"])
		payload := GetEventRequestPayload{
			ID: id,
		}
		if parser.ParseStructPayload(w, r, &payload) != true {
			return
		}
		var event models.Event = models.Event{ID: payload.ID}
		dbConn := db.GetDB()
		getRes := dbConn.Select("name", "description", "start_date", "end_date", "location").First(&event)
		if getRes.Error != nil {
			var errMsg string = "Error when getting event"
			var status int = 500
			if errors.Is(getRes.Error, gorm.ErrRecordNotFound) {
				errMsg = "Record not found"
				status = 404
			}
			responses.ErrResponse(w, &responses.ErrorResponse{
				Status:  status,
				Message: errMsg,
				Error:   nil,
			})
			logger.Errorf("Error when getting event: %s", getRes.Error.Error())
			return
		}
		res := GetEventResponsePayload{
			Name:        event.Name,
			Description: event.Description,
			StartDate:   event.StartDate,
			EndDate:     event.EndDate,
			Location:    event.Location,
		}
		responses.OkResponse(w, &responses.Response{
			Status:  200,
			Message: "Event found",
			Data:    res,
		})
	}
}

func UpdateEventHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, _ := uuid.Parse(vars["eventId"])
		payload := UpdateEventRequestPayload{
			ID: id,
		}
		if parser.ParseJSONPayload(w, r, &payload) != true {
			return
		}

		var event models.Event = models.Event{ID: payload.ID}
		dbConn := db.GetDB()
		getRes := dbConn.First(&event)
		if getRes.Error != nil {
			var errMsg string = "Error when getting event"
			var status int = 500
			if errors.Is(getRes.Error, gorm.ErrRecordNotFound) {
				errMsg = "Record not found"
				status = 404
			}
			responses.ErrResponse(w, &responses.ErrorResponse{
				Status:  status,
				Message: errMsg,
				Error:   nil,
			})
			logger.Errorf("Error when getting event: %s", getRes.Error.Error())
			return
		}
		event.Name = payload.Name
		event.Description = payload.Description
		event.StartDate = payload.StartDate
		event.EndDate = payload.EndDate
		event.Location = payload.Location
		dbConn.Select("name", "description", "start_date", "end_date", "location").Save(&event)

		res := UpdateEventResponsePayload{
			GetEventResponsePayload: GetEventResponsePayload{
				Name:        event.Name,
				Description: event.Description,
				StartDate:   event.StartDate,
				EndDate:     event.EndDate,
				Location:    event.Location,
			},
			ID: id,
		}
		responses.OkResponse(w, &responses.Response{
			Status:  200,
			Message: "Event updated",
			Data:    res,
		})
	}
}
