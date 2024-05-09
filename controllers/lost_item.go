package controllers

import (
	"github.com/arung-agamani/mitsukeru-server-go/responses"
	"github.com/arung-agamani/mitsukeru-server-go/services"
	"github.com/arung-agamani/mitsukeru-server-go/utils/parser"
	"github.com/arung-agamani/mitsukeru-server-go/utils/validator"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

type UpdateItemRequestPayload struct {
	ID              uuid.UUID `json:"id" validate:"required"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	ItemTypeID      string    `json:"itemTypeID"`
	ReporterContact string    `json:"reporterContact"`
	ReporterName    string    `json:"reporterName"`
	Returned        bool      `json:"returned"`
	Type            string    `json:"type"`
}

func CreateLostItemHandler(deps services.Dependencies) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload services.CreateLostItemPayload
		if parser.ParseJSONPayload(w, r, &payload) != true {
			return
		}
		lostItem, err := deps.LostItemService.CreateLostItem(payload)
		if err != nil {
			responses.HandleError(w, err)
		}
		responses.OkResponse(w, &responses.Response{
			Status:  201,
			Message: "LostItem has been created",
			Data:    lostItem,
		})
	}
}
func GetLostItemHandler(deps services.Dependencies) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, _ := uuid.Parse(vars["itemId"])
		if ok, err := validator.ValidateVariable(id, "required,uuid4"); !ok {
			responses.ErrResponse(w, &responses.ErrorResponse{
				Status:  400,
				Message: "invalid id",
				Error:   err,
			})
			return
		}
		item, err := deps.LostItemService.GetLostItem(id)
		if err != nil {
			responses.HandleError(w, err)
		}
		responses.OkResponse(w, &responses.Response{
			Status:  200,
			Message: "Retrieved Lost Item",
			Data:    item,
		})
		return
	}
}

func ListLostItemHandler(deps services.Dependencies) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lostItems, err := deps.LostItemService.ListLostItem()
		if err != nil {
			responses.HandleError(w, err)
			return
		}
		responses.OkResponse(w, &responses.Response{
			Status:  200,
			Message: "List of LostItem retrieved",
			Data:    lostItems,
		})
	}
}
func UpdateLostItemHandler(deps services.Dependencies) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, _ := uuid.Parse(vars["itemId"])
		payload := services.UpdateLostItemPayload{
			ID: id,
		}
		if parser.ParseJSONPayload(w, r, &payload) != true {
			return
		}
		item, err := deps.LostItemService.UpdateLostItem(payload)
		if err != nil {
			responses.HandleError(w, err)
			return
		}
		responses.OkResponse(w, &responses.Response{
			Status:  200,
			Message: "Updated",
			Data:    item,
		})

	}
}
func DeleteLostItemHandler(deps services.Dependencies) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
