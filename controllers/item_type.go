package controllers

import (
	"errors"
	"github.com/arung-agamani/mitsukeru-server-go/responses"
	"github.com/arung-agamani/mitsukeru-server-go/services"
	"github.com/arung-agamani/mitsukeru-server-go/utils/parser"
	"github.com/gorilla/mux"
	"net/http"
)

type CreateItemTypePayload struct {
	Name        string `json:"name" validate:"required,alphanum"`
	Description string `json:"description" validate:"required,max=256"`
}
type GetItemTypePayload struct {
	Name string `json:"name" validate:"required,alphanum"`
}

func CreateItemType(deps services.Dependencies) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload CreateItemTypePayload
		if parser.ParseJSONPayload(w, r, &payload) != true {
			return
		}
		itemType, err := deps.ItemTypeService.CreateItemType(payload.Name, payload.Description)
		if err != nil {
			if errors.Is(err, services.ValidatorError{}) {
				responses.ErrResponse(w, &responses.ErrorResponse{
					Status:  400,
					Message: "Bad request",
					Error:   err,
				})
			}
			responses.ErrResponse(w, &responses.ErrorResponse{
				Status:  500,
				Message: "Internal server error",
				Error:   err,
			})
			return
		}
		responses.OkResponse(w, &responses.Response{
			Status:  201,
			Message: "ItemType has been created",
			Data:    itemType,
		})
	}
}

func GetItemType(deps services.Dependencies) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name, _ := vars["name"]
		payload := GetItemTypePayload{Name: name}
		if parser.ParseStructPayload(w, r, &payload) != true {
			return
		}
		itemType, err := deps.ItemTypeService.GetItemType(payload.Name)
		if err != nil {
			if errors.Is(err, services.ValidatorError{}) {
				responses.ErrResponse(w, &responses.ErrorResponse{
					Status:  400,
					Message: "Bad request",
					Error:   err,
				})
			}
			responses.ErrResponse(w, &responses.ErrorResponse{
				Status:  500,
				Message: "Internal server error",
				Error:   err,
			})
			return
		}
		responses.OkResponse(w, &responses.Response{
			Status:  200,
			Message: "ItemType retrieved",
			Data:    itemType,
		})
	}
}
func ListItemType(deps services.Dependencies) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		itemTypes, err := deps.ItemTypeService.ListItemType()
		if err != nil {
			if errors.Is(err, services.ValidatorError{}) {
				responses.ErrResponse(w, &responses.ErrorResponse{
					Status:  400,
					Message: "Bad request",
					Error:   err,
				})
			}
			responses.ErrResponse(w, &responses.ErrorResponse{
				Status:  500,
				Message: "Internal server error",
				Error:   err,
			})
			return
		}
		responses.OkResponse(w, &responses.Response{
			Status:  200,
			Message: "List of ItemType retrieved",
			Data:    itemTypes,
		})
	}
}
func UpdateItemType(deps services.Dependencies) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
func DeleteItemType(deps services.Dependencies) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
