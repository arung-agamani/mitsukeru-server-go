package controllers

import (
	"errors"
	"github.com/arung-agamani/mitsukeru-server-go/db"
	"github.com/arung-agamani/mitsukeru-server-go/models"
	"github.com/arung-agamani/mitsukeru-server-go/responses"
	"github.com/arung-agamani/mitsukeru-server-go/utils/logger"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
)

func GetItemHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		itemId := vars["itemId"]
		dbConn := db.GetDB()
		logger.Infof("Request to item with id %s", itemId)
		var item models.LostItem
		err := dbConn.First(&item, itemId).Error
		if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
			res := &responses.ErrorResponse{
				Status:  404,
				Message: "item not found",
				Error:   nil,
			}
			responses.ErrResponse(w, res)
		} else {
			logger.Errorf("Unknown error: %v", err)
			responses.ErrResponse(w, &responses.ErrorResponse{
				Status:  500,
				Message: "Internal server error",
				Error:   nil,
			})
		}

		response := &responses.Response{
			Status:  200,
			Message: "success",
			Data:    item,
		}
		responses.OkResponse(w, response)
	}
}
